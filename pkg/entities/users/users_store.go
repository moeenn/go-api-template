package users

import (
	"app/core/database"
)

/**
 *  find a user against the provided ID
 *
 */
func GetUserByID(db *database.Database, id string) (User, error) {
	user := User{}
	if err := db.Conn.Get(&user, GET_USER_BY_ID_QUERY, id); err != nil {
		return User{}, err
	}

	return user, nil
}

/**
 *  add a new user to the database
 *
 */
func CreateUser(db *database.Database, user User) error {
	_, err := db.Conn.NamedExec(CREATE_USER_QUERY, user)
	if err != nil {
		return err
	}

	return nil
}
