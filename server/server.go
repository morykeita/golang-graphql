package main

import (
	"log"
	"net/http"
	"os"

	"github.com/morykeita/graphql-golang/database"

	graphql_golang "github.com/morykeita/graphql-golang"

	"github.com/99designs/gqlgen/handler"
)

const defaultPort = "8080"

func main() {

	DB := database.New(&pg.Options{
		User:     "",
		Password: "",
		Database: "",
	})

	DB.Close()
	DB.AddQueryHook(database.DBlogger{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	c := graphql_golang.Config{Resolvers: &graphql_golang.Resolver{}}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(graphql_golang.NewExecutableSchema(c)))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
