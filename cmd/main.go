package main

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/IcaroSilvaFK/gql-go/graph"
	"github.com/IcaroSilvaFK/gql-go/internal/database"
	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
)

const defaultPort = "8080"

func main() {

	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if !errors.Is(err, nil) {
		log.Fatal(err)
	}
	defer db.Close()

	port := os.Getenv("PORT")

	mx := chi.NewRouter()

	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{
					CategoryDB: database.NewCategory(db),
					CourseDB:   database.NewCourse(db),
				},
			},
		),
	)

	mx.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mx.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, mx))
}
