package repos

import (
	"example/graph/db"
	"example/internal/shared"

	"gorm.io/gorm"
)

func GetUserMap(con *gorm.DB, ids []uint) (map[string]db.User, error) {
	var users []db.User
	err := con.Find(&users, ids).Error
	if err != nil {
		return nil, err
	}

	res := make(map[string]db.User)
	for _, u := range users {
		res[*shared.EncodeCursor("users", u.ID)] = u
	}

	return res, nil
}
