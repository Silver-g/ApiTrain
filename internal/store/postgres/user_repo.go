package postgres

import (
	"ApiTrain/internal/domain"
)

func (r *Postgres) GetByUsername(username string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS (SELECT 1 FROM users WHERE username = $1)"
	err := r.db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (r *Postgres) Create(user *domain.User) (*domain.User, error) {
	query := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id"
	err := r.db.QueryRow(query, user.Username, user.Password).Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
