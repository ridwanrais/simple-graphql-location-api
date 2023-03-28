package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ridwanrais/simple-graphql-location-api/internal/presentation/graph"
)

const defaultPort = "8080"

func main() {
	// helloRepo := infrastructure.NewHelloRepository()
	// helloUseCase := usecase.HelloUseCase{}
	// resolver := graph.Resolver{
	// 	// HelloRepo:    helloRepo,
	// 	HelloUseCase: helloUseCase,
	// }

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver}))
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
