package Users

import (
	"fmt"
	db "gin-demo/modules/database/mysql"
)

type User struct {
	Id   int64
	Name string
}

func (User) TableName() string {
	return "users"
}

func GetField() []string {
	return []string{
		"id", "name",
	}
}

func GetOneById(id int64) (*User, error) {
	row := &User{}
	err := db.SqlDB.Select(GetField()).Where("id = ?", id).First(row).Error
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}
	return row, nil
}
