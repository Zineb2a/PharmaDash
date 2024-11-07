package util

import db "pharmaDashServer/db/sqlc"

type RegisterRequest struct {
	RegisterMode string
	UserData     *db.Account
}

type LoginRequest struct {
	Email    string
	Password string
}
