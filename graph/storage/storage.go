package storage

import (
	"context"
	"example/graph/db"
	"example/graph/model"
	"fmt"

	"github.com/graph-gophers/dataloader"
)

func GetUser(ctx context.Context, userIDs []string) ([]*model.User, error) {
	loaders := ForUser(ctx)
	thunk := loaders.LoadMany(ctx, dataloader.NewKeysFromStrings(userIDs))
	result, errors := thunk()
	if len(errors) > 0 {
		return nil, errors[0] // TODO: Join error
	}

	users := make([]*model.User, len(result))
	for i, user := range result {
		c := user.(db.User)
		users[i] = &model.User{
			ID:   fmt.Sprintf("%d", c.ID),
			Name: c.Name,
		}
	}

	return users, nil
}
