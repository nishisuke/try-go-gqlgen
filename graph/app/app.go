package app

import (
	"context"
	"example/graph/model"
	"fmt"
)

type (
	ctxKey string
	Foo    interface {
		GetUser(context.Context, []string) ([]*model.User, error)
	}

	App struct {
	}
)

const (
	key = ctxKey("loader")
)

func For(ctx context.Context) Foo {
	return ctx.Value(key).(Foo)
}

func StoreLoader(ctx context.Context, l Foo) context.Context {
	return context.WithValue(ctx, key, l)
}

func NewApp() App {
	return App{}
}

func (a App) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {

	panic(fmt.Errorf("not implemented"))
}
func (a App) Todos(ctx context.Context) ([]*model.Todo, error) {

	friend := model.User{
		ID:   "fri",
		Name: "friend",
	}
	user := model.User{
		ID:      "userid",
		Name:    "hey",
		Friends: []*model.User{&friend, &friend},
	}
	return []*model.Todo{
		{ID: "2", Text: "Wash", Done: false, User: &user},
		{ID: "3", Text: "Clean", Done: false, User: &user},
		{ID: "4", Text: "Eat", Done: true, User: &user},
	}, nil
}

func (a App) Friends(ctx context.Context, obj *model.User) ([]*model.User, error) {
	friendIDs := []string{"1", "2"}
	load := For(ctx)
	return load.GetUser(ctx, friendIDs)
}
