package storage

// import graph gophers with your other imports
import (
	"context"
	"errors"
	"example/graph/db"

	"github.com/graph-gophers/dataloader"
)

type (
	IGetUsers interface {
		GetUsers(context.Context, []string) (map[string]db.User, error)
	}
)

func NewUsersLoader(l IGetUsers) *dataloader.Loader {
	u := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		um, err := l.GetUsers(ctx, conv(keys))

		if err != nil {
			return []*dataloader.Result{{Error: err}}
		}

		return packResult(keys, um)
	}

	return dataloader.NewBatchedLoader(u)
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
