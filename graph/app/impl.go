package app

import (
	"context"
	"example/graph/loader"
	"example/graph/model"
	"fmt"

	"gorm.io/gorm"
)

const key = "loader-key"

var (
	arr = []*model.Todo{
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
	}
)

func StoreLoader(ctx context.Context, con *gorm.DB) context.Context {
	return context.WithValue(ctx, key, loader.NewLoader(con))
}

func QueryTodos(ctx context.Context, first *int, after *string) (*model.TodoConnection, error) {
	edges := make([]*model.TodoEdge, len(arr))
	for i, a := range arr {
		edges[i] = &model.TodoEdge{
			Node: a, Cursor: a.ID,
		}
	}
	con := model.TodoConnection{
		Edges:    edges,
		PageInfo: &model.PageInfo{},
	}
	return &con, nil
}

func CreateTodo(ctx context.Context, input model.NewTodo) (*model.TodoEdge, error) {
	todo := &model.Todo{
		ID:   fmt.Sprintf("%s1", arr[len(arr)-1].ID),
		Text: input.Text,
		Done: false,
		User: &model.User{ID: "usera", Name: "usera"}}

	arr = append(arr, todo)
	return &model.TodoEdge{Node: todo}, nil
}
