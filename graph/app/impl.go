package app

import (
	"context"
	"example/graph/model"
	"fmt"
)

func CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
	return nil, nil
}
func QueryTodos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
	return nil, nil
}
func QueryFriends(ctx context.Context, obj *model.User) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
	return nil, nil
}
