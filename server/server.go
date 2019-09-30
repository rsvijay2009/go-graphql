package main

import (
	"log"
	"net/http"
	"os"

	go_graphql "github.com/rsvijay2009/go-graphql"

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
