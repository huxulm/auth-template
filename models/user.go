package models

import "errors"

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"-"`
}

var ErrUserNotExists = errors.New("user not exists")

// user mapping
var db = map[string]*User{
	"foo": {"1", "foo", "foo123"},
	"bar": {"2", "bar", "bar123"},
}

func FindUserByName(name string) (*User, error) {
	if v, ok := db[name]; ok {
		return v, nil
	}
	return nil, ErrUserNotExists
}
