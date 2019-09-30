package main

import (
	go_graphql "go-graphql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
)

func main() {
	go_graphql.ConnectDB()

	port := os.Getenv("PORT")
	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(go_graphql.NewExecutableSchema(go_graphql.Config{Resolvers: &go_graphql.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
