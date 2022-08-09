package usecases

import (
	"context"
	"example/graph/db"
	"example/graph/model"
	"example/internal/shared"
	"fmt"

	"gorm.io/gorm"
)

type (
	TodoConnectionUsecase struct{}
)

const (
	dbkey = "gorm-key"
)

const (
	defaultLimit = 10
)

func limit(c *int) int {
	if c != nil {
		return *c
	} else {
		return defaultLimit
	}
}

func (u TodoConnectionUsecase) Fetch(ctx context.Context, first *int, after *string, scope func(*gorm.DB) *gorm.DB) (*model.TodoConnection, error) {
	scpFn := scope
	if scpFn == nil {
		scpFn = func(d *gorm.DB) *gorm.DB { return d }
	}

	con := ctx.Value(dbkey).(*gorm.DB)
	limit := limit(first)

	var res []db.Todo
	if after == nil {
		err := con.Model(&db.Todo{}).Scopes(scpFn).Order("id asc").Limit(limit).Find(&res).Error

		if err != nil {
			return nil, err
		}
	} else {
		decoded, err := shared.DecodeCursor(after)
		if err != nil {
			return nil, err
		}
		err = con.Model(&db.Todo{}).Scopes(scpFn).Where("id > ?", decoded).Order("id asc").Limit(limit).Find(&res).Error
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
	err := con.Model(&db.Todo{}).Scopes(scpFn).Where("id < ?", res[0].ID).Order("id asc").Limit(1).Find(&hasPre).Error
	if err != nil {
		return nil, err
	}
	err = con.Model(&db.Todo{}).Scopes(scpFn).Where("id > ?", res[len(res)-1].ID).Order("id asc").Limit(1).Find(&hasNex).Error
	if err != nil {
		return nil, err
	}

	return &model.TodoConnection{
		Edges: edges,
		PageInfo: &model.PageInfo{
			EndCursor:       shared.EncodeCursor("todos", res[len(res)-1].ID),
			HasNextPage:     len(hasNex) > 0,
			HasPreviousPage: len(hasPre) > 0,
			StartCursor:     shared.EncodeCursor("todos", res[0].ID),
		},
	}, nil
}

func todoEdge(todo db.Todo) *model.TodoEdge {
	cur := shared.EncodeCursor("todos", todo.ID)
	return &model.TodoEdge{
		Node: &model.Todo{
			ID:     *cur,
			Text:   fmt.Sprintf("%d %s", todo.ID, todo.Text),
			Done:   todo.Done,
			UserID: *shared.EncodeCursor("users", todo.UserID),
		}, Cursor: *cur,
	}
}
