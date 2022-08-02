package app

import (
	"context"
	"example/graph/loader"
	"example/graph/model"
	"fmt"

	"gorm.io/gorm"
)

const key = "loader-key"

func StoreLoader(ctx context.Context, con *gorm.DB) context.Context {
	return context.WithValue(ctx, key, loader.NewLoader(con))
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
	friendIDs := []string{"1", "2"} // select user_id from friends where from_id = obj.id
	return loader.GetUsers(ctx, friendIDs)
}

func QueryUsers(ctx context.Context) ([]*model.User, error) {
	loader := ctx.Value(key).(*loader.Loader)
	allIDs := []string{"1", "2"}
	return loader.GetUsers(ctx, allIDs)
}
