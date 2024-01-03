package data

import (
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type UserModel struct {
	DB *pgxpool.Pool
}

type Models struct {
	Users UserModel
}

func NewModels(db *pgxpool.Pool) Models {
	return Models{
		Users: UserModel{DB: db},
	}
}
