package database_sqlite

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
)

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
