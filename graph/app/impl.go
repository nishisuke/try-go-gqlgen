package app

import (
	"context"
	"example/graph/loader"
	"example/graph/model"
	"fmt"
)

const key = "loader-key"

func StoreLoader(ctx context.Context, loader *loader.Loader) context.Context {
	return context.WithValue(ctx, key, loader)
}

func CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
	return nil, nil
}
func QueryTodos(ctx context.Context) ([]*model.Todo, error) {
	return []*model.Todo{
		{User: &model.User{}},
	}, nil
}
func QueryFriends(ctx context.Context, obj *model.User) ([]*model.User, error) {
	loader := ctx.Value(key).(*loader.Loader)
	return loader.GetUser(ctx, []string{"1", "2"}) // select user_id from friends where from_id = obj.id
}
