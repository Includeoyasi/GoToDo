package repository

import (
	"github.com/Includeoyasi/todo-app"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user todo.User) (int, error) {
	return 0, nil
}
