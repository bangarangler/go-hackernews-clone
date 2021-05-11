package main

import (
	// "database/sql"
	"log"
	"net/http"
	"os"

	// "github.com/bangarangler/go-hackernews-clone/internal/pkg/pg"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/bangarangler/go-hackernews-clone/graph"
	"github.com/bangarangler/go-hackernews-clone/graph/generated"
)

const defaultPort = "8080"

func main() {
	// port := os.Getenv("PORT")
	port := database.goDotEnvVar("PORT").Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	database.InitDB()
	database.Migrate()
	server := handler.GraphQL(go_hackernews_clone.NewExecutableSchema(go_hackernews_clone.Config(Resolvers: &go_hackernews_clone.Resolver())))
	router.Handler("/", handler.Playground("GraphQL playground", "/query"))
	router.Handler("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

	// srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
  //
	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", srv)
  //
	// log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// log.Fatal(http.ListenAndServe(":"+port, nil))
}
