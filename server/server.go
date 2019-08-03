package main

import (
	"github.com/graphql-gopher/database"
	"log"
	"net/http"
	"os"
	"github.com/99designs/gqlgen/handler"
	"github.com/graphql-gopher"
)

const defaultPort = "8080"

func main() {
	err, _ := database.SingleDBMongo()
	if err != nil {
		panic(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(graphql_gopher.NewExecutableSchema(graphql_gopher.Config{Resolvers: &graphql_gopher.Resolver{}})))
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
