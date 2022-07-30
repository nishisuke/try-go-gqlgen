package storage

// import graph gophers with your other imports
import (
	"context"
	"errors"
	"example/graph/db"
	"fmt"

	"github.com/graph-gophers/dataloader"
	"gorm.io/gorm"
)

type (
	Foo struct {
		Conn func(context.Context) *gorm.DB
	}
	IGetUsers interface {
		GetUsers(context.Context, dataloader.Keys) []*dataloader.Result
	}
)

func NewFoo(con func(context.Context) *gorm.DB) *Foo {
	return &Foo{con}
}

func NewUsersLoader(l IGetUsers) *dataloader.Loader {
	return dataloader.NewBatchedLoader(l.GetUsers)
}

func (l *Foo) GetUsers(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	con := l.Conn(ctx)
	um, err := aa(con, keys)

	if err != nil {
		return []*dataloader.Result{
			{Error: err},
		}
	}

	res := make([]*dataloader.Result, len(keys))

	for i, key := range keys {
		if u, ok := um[key.String()]; ok {
			res[i] = &dataloader.Result{Data: u}
		} else {
			res[i] = &dataloader.Result{Error: errors.New("NotFound")}
		}
	}

	return res
}

func aa(d *gorm.DB, keys dataloader.Keys) (map[string]db.User, error) {
	var users []db.User
	if err := d.Find(&users, keys).Error; err != nil {
		return nil, err
	}

	res := make(map[string]db.User)

	for _, u := range users {
		res[fmt.Sprintf("%d", u.ID)] = u
	}

	return res, nil
}
