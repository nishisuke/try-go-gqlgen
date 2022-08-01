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
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{App: app.NewApp()}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", Middleware(con, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func NewLoader(con *gorm.DB) *storage.Loader {
	useRepo := repos.NewUserRepo(func(ctx context.Context) *gorm.DB {
		return con.Debug().WithContext(ctx)
	})
	return storage.NewLoader(useRepo)
}
func Middleware(con *gorm.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCtx := app.StoreLoader(r.Context(), NewLoader(con))
		r = r.WithContext(nextCtx)
		next.ServeHTTP(w, r)
	})
}
