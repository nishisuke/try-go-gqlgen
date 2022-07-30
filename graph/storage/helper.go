package storage

// import graph gophers with your other imports
import (
	"context"
	"net/http"

	"github.com/graph-gophers/dataloader"
)

type (
	ctxKey string
)

const (
	userloadersKey = ctxKey("userKey")
)

func Middleware(userloaders *dataloader.Loader, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCtx := context.WithValue(r.Context(), userloadersKey, userloaders)
		r = r.WithContext(nextCtx)
		next.ServeHTTP(w, r)
	})
}

func ForUser(ctx context.Context) *dataloader.Loader {
	return ctx.Value(userloadersKey).(*dataloader.Loader)
}
