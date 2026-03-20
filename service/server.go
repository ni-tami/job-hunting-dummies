package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	migration "github.com/ni-tami/job-hunting-dummies-service/db"
	"github.com/ni-tami/job-hunting-dummies-service/internal/client"
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

const (
	defaultPort = "8888"
 	grpcPort = 50051
)

func main() {
	migration.MigrateTables()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:8080",
			"http://localhost:3000",
		},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)
	// router.Use(auth.Middleware())
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
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
