package repos

import (
	"example/graph/db"
	"fmt"

	"gorm.io/gorm"
)

func GetUserMap(con *gorm.DB, ids []string) (map[string]db.User, error) {
	var users []db.User
	err := con.Find(&users, ids).Error
	if err != nil {
		return nil, err
	}

	res := make(map[string]db.User)
	for _, u := range users {
		res[fmt.Sprintf("%d", u.ID)] = u
	}

	return res, nil
}
