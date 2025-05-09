package postgres

import (
	"ApiTrain/internal/domain"
	"database/sql"
)

type PostgresUserRepo struct {
	db *sql.DB
}

func NewPostgresUserRepo(dataBase *sql.DB) *PostgresUserRepo {
	var userRepoPointer PostgresUserRepo
	userRepoPointer.db = dataBase
	return &userRepoPointer
}

func (r *PostgresUserRepo) Create(user *domain.User) (*domain.User, error) {
	query := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id"
	err := r.db.QueryRow(query, user.Username, user.Password).Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
