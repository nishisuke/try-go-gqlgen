package storage

// import graph gophers with your other imports
import (
	"context"
	"errors"
	"example/graph/model"
	"net/http"

	"github.com/graph-gophers/dataloader"
)

type (
	ctxKey string
	Foo    interface {
		GetUser(context.Context, dataloader.Keys) ([]*model.User, error)
	}
)

const (
	key = ctxKey("loader")
)

func Middleware(loaders Foo, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCtx := context.WithValue(r.Context(), key, loaders)
		r = r.WithContext(nextCtx)
		next.ServeHTTP(w, r)
	})
}

func For(ctx context.Context) Foo {
	return ctx.Value(key).(Foo)
}

func conv(keys dataloader.Keys) []string {
	ids := make([]string, len(keys))
	for i, k := range keys {
		ids[i] = k.String()
	}
	return ids
}

func packResult[T any](keys dataloader.Keys, look map[string]T) []*dataloader.Result {

	res := make([]*dataloader.Result, len(keys))

	for i, key := range keys {
		if u, ok := look[key.String()]; ok {
			res[i] = &dataloader.Result{Data: u}
		} else {
			res[i] = &dataloader.Result{Error: errors.New("NotFound")}
		}
	}

	return res
}

func GetUser(ctx context.Context, userIDs []string) ([]*model.User, error) {
	loaders := For(ctx)
	return loaders.GetUser(ctx, dataloader.NewKeysFromStrings(userIDs))
}
