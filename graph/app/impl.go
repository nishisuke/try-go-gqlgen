package app

import (
	"context"
	"example/graph/db"
	"example/graph/model"
	"fmt"

	"gorm.io/gorm"
)

const (
	defaultLimit = 10
)

func QueryTodos(ctx context.Context, first *int, after *string) (*model.TodoConnection, error) {
	con := ctx.Value(dbkey).(*gorm.DB)

	limit := defaultLimit
	if first != nil {
		limit = *first
	}

	var res []db.Todo
	if after == nil {
		err := con.Model(&db.Todo{}).Order("id asc").Limit(limit).Find(&res).Error

		if err != nil {
			return nil, err
		}
	} else {
		decoded, err := decodeCursor(after)
		if err != nil {
			return nil, err
		}
		err = con.Model(&db.Todo{}).Where("id > ?", decoded).Order("id asc").Limit(limit).Find(&res).Error
		if err != nil {
			return nil, err
		}
	}
	if len(res) == 0 {
		p, err := emptyTodoPageInfo(con, after)
		if err != nil {
			return nil, err
		}
		return &model.TodoConnection{PageInfo: p}, nil
	}

	edges := make([]*model.TodoEdge, len(res))
	for i, todo := range res {
		edges[i] = todoEdge(todo)
	}

	var hasPre []db.Todo
	var hasNex []db.Todo
	err := con.Model(&db.Todo{}).Where("id < ?", res[0].ID).Order("id asc").Limit(1).Find(&hasPre).Error
	if err != nil {
		return nil, err
	}
	err = con.Model(&db.Todo{}).Where("id > ?", res[len(res)-1].ID).Order("id asc").Limit(1).Find(&hasNex).Error
	if err != nil {
		return nil, err
	}

	return &model.TodoConnection{
		Edges: edges,
		PageInfo: &model.PageInfo{
			EndCursor:       encodeCursor("todos", res[len(res)-1].ID),
			HasNextPage:     len(hasNex) > 0,
			HasPreviousPage: len(hasPre) > 0,
			StartCursor:     encodeCursor("todos", res[0].ID),
		},
	}, nil
}

func emptyTodoPageInfo(con *gorm.DB, after *string) (*model.PageInfo, error) {
	pre := false
	if after != nil {
		decoded, err := decodeCursor(after)
		if err != nil {
			return nil, err
		}
		var has []db.Todo
		err = con.Model(&db.Todo{}).Where("id <= ?", decoded).Order("id asc").Limit(1).Find(&has).Error
		if err != nil {
			return nil, err
		}

		pre = len(has) > 0
	}
	// TODO: before -> hasNext
	nex := true
	return &model.PageInfo{
		HasNextPage:     nex,
		HasPreviousPage: pre,
	}, nil
}

func todoEdge(todo db.Todo) *model.TodoEdge {
	cur := encodeCursor("todos", todo.ID)
	return &model.TodoEdge{
		Node: &model.Todo{
			ID:   *cur,
			Text: fmt.Sprintf("%d %s", todo.ID, todo.Text),
			Done: todo.Done,
		}, Cursor: *cur,
	}
}

func CreateTodo(ctx context.Context, input model.NewTodo) (*model.TodoEdge, error) {
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
