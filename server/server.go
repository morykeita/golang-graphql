package main

import (
	"github.com/go-pg/pg/v9"
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
		User:     "rxcrztvs",
		Password: "GrwFnC2mhTYL7zys4KmTOgNE6v5AGfjX",
		Database: "rxcrztvs",
		Addr:"rajje.db.elephantsql.com:5432",
	})
	//TODO  why is connection  closing before query is executed ? sslmode maybe ?
	//DB.Close()
	DB.AddQueryHook(database.DBlogger{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	c := graphql_golang.Config{Resolvers: &graphql_golang.Resolver{
		MeetupsRepository : database.MeetupsRepository{DB:DB},
		UserRepository :database.UserRepository{DB:DB},
	}}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(graphql_golang.NewExecutableSchema(c)))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
