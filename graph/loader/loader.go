package loader

// import graph gophers with your other imports
import (
	"context"
	"errors"
	"example/graph/db"
	"example/graph/model"
	"example/graph/repos"
	"example/internal/shared"
	"fmt"

	"github.com/graph-gophers/dataloader"
	"gorm.io/gorm"
)

type (
	Loader struct {
		UserLoader *dataloader.Loader
	}
)

func NewLoader(con *gorm.DB) *Loader {
	return &Loader{
		UserLoader: newUserLoader(con),
	}
}

func (l *Loader) GetUser(ctx context.Context, userID string) (*model.User, error) {
	thunk := l.UserLoader.Load(ctx, dataloader.StringKey(userID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}

	c := result.(db.User)
	return &model.User{
		ID:   *shared.EncodeCursor("users", c.ID),
		Name: c.Name,
	}, nil
}
func (l *Loader) GetUsers(ctx context.Context, userIDs []string) ([]*model.User, error) {
	thunk := l.UserLoader.LoadMany(ctx, dataloader.NewKeysFromStrings(userIDs))
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

func keysToStrings(keys dataloader.Keys) []string {
	ids := make([]string, len(keys))
	for i, k := range keys {
		ids[i] = k.String()
	}
	return ids
}

func newUserLoader(con *gorm.DB) *dataloader.Loader {
	u := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]uint, len(keys))
		for i, k := range keys {
			c := k.String()
			id, _ := shared.DecodeCursor(&c)
			// TODO: handle err
			ids[i] = uint(id)
		}

		lookup, err := repos.GetUserMap(con.WithContext(ctx).Debug(), ids)

		if err != nil {
			return []*dataloader.Result{{Error: err}}
		}

		res := make([]*dataloader.Result, len(keys))

		for i, key := range keys {
			if u, ok := lookup[key.String()]; ok {
				res[i] = &dataloader.Result{Data: u}
			} else {
				res[i] = &dataloader.Result{Error: errors.New("NotFound")}
			}
		}

		return res
	}
	return dataloader.NewBatchedLoader(u)
}
