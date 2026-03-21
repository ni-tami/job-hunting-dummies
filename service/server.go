package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	migration "github.com/ni-tami/job-hunting-dummies-service/db"
	"github.com/ni-tami/job-hunting-dummies-service/internal/client"
	"github.com/ni-tami/job-hunting-dummies-service/internal/config"
	"github.com/ni-tami/job-hunting-dummies-service/internal/dataloader"
	"github.com/ni-tami/job-hunting-dummies-service/internal/graphql"
	"github.com/ni-tami/job-hunting-dummies-service/internal/graphql/model"
	grpc_service "github.com/ni-tami/job-hunting-dummies-service/internal/grpc/service"
	"github.com/ni-tami/job-hunting-dummies-service/internal/repository"
	"github.com/ni-tami/job-hunting-dummies-service/internal/usecase"
	pb "github.com/ni-tami/job-hunting-dummies-service/pb/out/go/job_hunting_dummies"
	"github.com/rs/cors"
	"github.com/vektah/gqlparser/v2/ast"
	"google.golang.org/grpc"
)

func main() {
	config.Load()
	migration.MigrateTables()
	gqlPort := config.Koanf.Int("gql.port")
	grpcPort := config.Koanf.Int("grpc.port")

	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   config.Koanf.MustStrings("cors.allowed_origins"),
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	db := client.NewCrdbConn()
	repo := repository.NewJobPortalRepository(db)
	usecase := usecase.NewJobPortalUsecase(repo)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	jobHuntService := grpc_service.NewJobHuntServerImpl(usecase)
	pb.RegisterJobHuntServiceServer(grpcServer, jobHuntService)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("grpc server failed to serve: %v", err)
		}
	}()

	// GQL Server
	srv := handler.New(model.NewExecutableSchema(
		model.Config{
			Resolvers: graphql.NewResolver(usecase),
		}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", dataloader.Middleware(db, srv))

	log.Printf("connect to http://localhost:%d/ for grpc", grpcPort)
	log.Printf("connect to http://localhost:%d/ for GraphQL playground", gqlPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", gqlPort), router))
}
