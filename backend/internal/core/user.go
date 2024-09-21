package core

import "context"

type (
	User struct {
		ID int `gorm:"column:id"`
	}

	UserStore interface {
		AddUser(ctx context.Context, user User) (id int, err error)
	}
)

func (User) TableName() string {
	return "users"
}
