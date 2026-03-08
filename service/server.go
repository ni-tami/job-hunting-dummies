package main

import (
	"log"
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
	"github.com/ni-tami/job-hunting-dummies-service/internal/repository"
	"github.com/ni-tami/job-hunting-dummies-service/internal/usecase"
	"github.com/rs/cors"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8888"

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
