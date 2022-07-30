package main

import (
	"context"
	"example/graph"
	"example/graph/app"
	"example/graph/db"
	"example/graph/generated"
	"example/graph/repos"
	"example/graph/storage"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"gorm.io/gorm"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// DB
	con, err := db.Setup()
	if err != nil {
		log.Fatalln(err)
	}
	con.AutoMigrate(&db.User{})

	// Graph
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{App: app.App{}}}))

	useRepo := repos.NewUserRepo(func(ctx context.Context) *gorm.DB {
		return con.Debug().WithContext(ctx)
	})
	l := storage.NewLoader(useRepo)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", storage.Middleware(l, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
