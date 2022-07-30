package graph

import (
	"context"
	"example/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	App IApp
}

type IApp interface {
	CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error)
	Todos(ctx context.Context) ([]*model.Todo, error)
	Friends(ctx context.Context, obj *model.User) ([]*model.User, error)
}
