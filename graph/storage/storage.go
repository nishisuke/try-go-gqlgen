package storage

// import graph gophers with your other imports
import (
	"context"
	"example/graph/model"
	"fmt"

	"github.com/graph-gophers/dataloader"
)

// GetUser wraps the User dataloader for efficient retrieval by user ID
func GetUser(ctx context.Context, userIDs []string) ([]*model.User, error) {
	loaders := For(ctx)
	thunk := loaders.UserLoader.LoadMany(ctx, dataloader.NewKeysFromStrings(userIDs))
	result, errors := thunk()
	fmt.Println(errors)
	if len(errors) > 0 {
		return nil, errors[0]
	}

	users := make([]*model.User, len(result))
	for i, user := range result {
		users[i] = user.(*model.User)
	}

	return users, nil
}
