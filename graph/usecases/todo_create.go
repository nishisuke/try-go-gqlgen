package usecases

import (
	"context"
	"example/graph/db"
	"example/graph/model"

	"gorm.io/gorm"
)

type (
	TodoCreateUsecase struct {
	}
)

func (u TodoCreateUsecase) Create(ctx context.Context, input model.NewTodo) (*model.TodoEdge, error) {
	todo := &db.Todo{
		Text:   input.Text,
		Done:   false,
		UserID: 1,
	}

	con := ctx.Value(dbkey).(*gorm.DB)
	err := con.Create(todo).Error
	if err != nil {
		return nil, err
	}

	return todoEdge(*todo), nil
}
