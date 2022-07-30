package storage

// import graph gophers with your other imports
import (
	"context"
	"example/graph/db"
	"example/graph/model"
	"fmt"

	"github.com/graph-gophers/dataloader"
)

type (
	IGetUsers interface {
		GetUsers(context.Context, []string) (map[string]db.User, error)
	}
	Loader struct {
		UserLoader *dataloader.Loader
	}
)

func (l *Loader) GetUser(ctx context.Context, keys dataloader.Keys) ([]*model.User, error) {
	thunk := l.UserLoader.LoadMany(ctx, keys)
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

func NewLoader(l IGetUsers) *Loader {
	u := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		um, err := l.GetUsers(ctx, conv(keys))

		if err != nil {
			return []*dataloader.Result{{Error: err}}
		}

		return packResult(keys, um)
	}

	return &Loader{
		UserLoader: dataloader.NewBatchedLoader(u),
	}
}
