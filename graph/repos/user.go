package repos

// import graph gophers with your other imports
import (
	"context"
	"example/graph/db"
	"fmt"

	"gorm.io/gorm"
)

type (
	UserRepo struct {
		Conn func(context.Context) *gorm.DB
	}
)

func NewUserRepo(con func(context.Context) *gorm.DB) UserRepo {
	return UserRepo{con}
}

func (l UserRepo) GetUsers(ctx context.Context, keys []string) (map[string]db.User, error) {
	con := l.Conn(ctx)

	var users []db.User
	if err := con.Find(&users, keys).Error; err != nil {
		return nil, err
	}

	res := make(map[string]db.User)

	for _, u := range users {
		res[fmt.Sprintf("%d", u.ID)] = u
	}

	return res, nil
}
