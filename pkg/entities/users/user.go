package users

import (
	"app/core/helpers"
)

type User struct {
	Id       string `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

/* create instance of a User (without db persistence */
func NewUser(name, email, passwordHash string) User {
	return User{
		Id:       helpers.UUID(),
		Name:     name,
		Email:    email,
		Password: passwordHash,
	}
}
