package app

import (
	"context"
	"example/graph/loader"
	"example/graph/model"
	"example/graph/usecases"
	"example/internal/shared"

	"gorm.io/gorm"
)

func QueryTodos(ctx context.Context, first *int, after *string) (*model.TodoConnection, error) {
	u := usecases.TodoConnectionUsecase{}
	return u.Fetch(ctx, first, after, nil)
}

func CreateTodo(ctx context.Context, input model.NewTodo) (*model.TodoEdge, error) {
	return usecases.TodoCreateUsecase{}.Create(ctx, input)
}

func User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	load := ctx.Value(key).(*loader.Loader)
	return load.GetUser(ctx, obj.UserID)
}

func Todos(ctx context.Context, obj *model.User, first *int, after *string) (*model.TodoConnection, error) {
	decoded, err := shared.DecodeCursor(&obj.ID)
	if err != nil {
		return nil, err
	}
	u := usecases.TodoConnectionUsecase{}
	return u.Fetch(ctx, first, after, func(d *gorm.DB) *gorm.DB {
		return d.Where("user_id = ?", decoded)
	})
}
func TotalCount(ctx context.Context, obj *model.TodoConnection) (int, error) {
	con := ctx.Value(dbkey).(*gorm.DB)

	var count int64
	err := con.Model(&model.Todo{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}
