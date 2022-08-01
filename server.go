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
	"gorm.io/gorm"
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

	http.Handle("/query", middleware(con, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func middleware(con *gorm.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCtx := app.StoreLoader(r.Context(), loader.NewLoader(con))
		next.ServeHTTP(w, r.WithContext(nextCtx))
	})
}
