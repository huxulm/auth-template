package services

import (
	"errors"

	db "github.com/huxulm/auth-template/models"
)

// Authenticate by simple {name, password} pair
func Authenticate(name, pass string) (*db.User, error) {
	if u, err := db.FindUserByName(name); err != nil {
		return nil, err
	} else {
		if u.Password != pass {
			return nil, errors.New("invalid password")
		}
		return u, nil
	}
}
