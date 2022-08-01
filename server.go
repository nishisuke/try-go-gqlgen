package main

import (
	"example/graph"
	"example/graph/app"
	"example/graph/db"
	"example/graph/generated"
	"example/graph/loader"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	// DB
	con, err := db.Setup()
	if err != nil {
		log.Fatalln(err)
	}
	con.AutoMigrate(&db.User{})
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	load := loader.NewLoader(con)
	http.Handle("/query", middleware(load, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func middleware(loader *loader.Loader, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCtx := app.StoreLoader(r.Context(), loader)
		next.ServeHTTP(w, r.WithContext(nextCtx))
	})
}
