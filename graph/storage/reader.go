package storage

import (
	"context"
	"database/sql"
	"example/graph/model"

	"github.com/graph-gophers/dataloader"
)

type (
	UserReader struct {
		conn *sql.DB
	}
)

func (u *UserReader) GetUsers(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	res := make([]*dataloader.Result, len(keys))
	for i, key := range keys {
		user := &model.User{ID: key.String(), Name: "name"}
		res[i] = &dataloader.Result{Data: user}
	}

	return res
}
