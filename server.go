package main

import (
	"example/depth"
	"example/graph"
	"example/graph/app"
	"example/graph/db"
	"example/graph/generated"
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
	con.AutoMigrate(
		&db.User{},
		&db.Todo{},
	)

	// Port
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	conf := generated.Config{Resolvers: &graph.Resolver{}}
	//conf.Complexity.Query.Todos = func(childComplexity int, first *int, after *string) int {
	//	if first == nil {
	//		return childComplexity // TODO: Adjust
	//	} else {
	//		return childComplexity * *first // TODO: Adjust
	//	}
	//}
	//conf.Complexity.Todo.User = func(childComplexity int) int {
	//	return childComplexity * 2 // TODO: Adjust
	//}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(conf))
	// srv.Use(extension.FixedComplexityLimit(850)) // Adjust
	srv.Use(depth.NewFixedMaxDepthLimit(5))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", middleware(con, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func middleware(con *gorm.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCtx := app.StoreLoader(r.Context(), con)
		nextCtx = app.StoreDB(nextCtx, con)
		next.ServeHTTP(w, r.WithContext(nextCtx))
	})
}
