package app

import (
	"context"
	"example/graph/loader"
	"example/graph/model"

	"gorm.io/gorm"
)

const key = "loader-key"

func StoreLoader(ctx context.Context, con *gorm.DB) context.Context {
	return context.WithValue(ctx, key, loader.NewLoader(con))
}
func QueryTodos(ctx context.Context) ([]*model.Todo, error) {
	return []*model.Todo{
		{
			ID:   "ida",
			Text: "hey",
			Done: false,
			User: &model.User{ID: "usera", Name: "usera"}},
		{
			ID:   "idb",
			Text: "foo",
			Done: false,
			User: &model.User{ID: "usera", Name: "usera"}},
		{
			ID:   "idc",
			Text: "bar",
			Done: false,
			User: &model.User{ID: "usera", Name: "usera"}},
	}, nil
}
