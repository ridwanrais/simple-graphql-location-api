package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/ridwanrais/simple-graphql-location-api/internal/presentation/graph"
	"github.com/sirupsen/logrus"
)

const defaultPort = "8080"

func main() {
	if err := godotenv.Load(); err != nil {
		logrus.Warningln("Can't find env.file. The app is now run using default environment variables.")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	resolver := graph.InitResolver()
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
