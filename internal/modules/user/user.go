package user

import "database/sql"

type UserRole string

type User struct {
	Id       string         `db:"user_id"`
	Email    string         `db:"email"`
	Password sql.NullString `db:"password"`
	Role     UserRole       `db:"role"`
	IsActive bool           `db:"is_active"`
}
