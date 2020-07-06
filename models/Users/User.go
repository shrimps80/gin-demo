package Users

import (
	"fmt"
	db "gin-demo/modules/database/mysql"
)

type User struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (User) TableName() string {
	return "users"
}

func GetField() []string {
	return []string{
		"id", "name", "email",
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

func GetOneByEmail(email string) (*User, error) {
	row := &User{}
	err := db.SqlDB.Select(GetField()).Where("email = ?", email).First(row).Error
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}
	return row, nil
}
