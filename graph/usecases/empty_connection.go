package usecases

import (
	"example/graph/db"
	"example/graph/model"
	"example/internal/shared"

	"gorm.io/gorm"
)

func emptyTodoPageInfo(con *gorm.DB, after *string) (*model.PageInfo, error) {
	pre := false
	if after != nil {
		decoded, err := shared.DecodeCursor(after)
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
