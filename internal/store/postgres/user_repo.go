package postgres

import (
	"ApiTrain/internal/domain"
	"database/sql"
	"errors"
)

var ErrUserNotFound = errors.New("user not found") // это потом убрать в отдельное место

func (r *Postgres) GetByUsername(username string) (*domain.User, error) {
	var user domain.User
	query := "SELECT id, username, password FROM users WHERE username = $1"
	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password) // почистить значения которые вовращает функция как минимум пароль тут не нужен точно(+ тести лох)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *Postgres) Create(user *domain.User) (*domain.User, error) {
	query := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id"
	err := r.db.QueryRow(query, user.Username, user.Password).Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
