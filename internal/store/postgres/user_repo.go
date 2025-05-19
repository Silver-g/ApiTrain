package postgres

import (
	"ApiTrain/internal/domain"
	"database/sql"
	"errors"
)

type PostgresUserRepo struct {
	db *sql.DB
}

func NewPostgresUserRepo(dataBase *sql.DB) *PostgresUserRepo {
	var userRepoPointer PostgresUserRepo
	userRepoPointer.db = dataBase
	return &userRepoPointer
}

var ErrUserNotFound = errors.New("user not found") // это потом убрать в отдельное место

func (r *PostgresUserRepo) GetByUsername(username string) (*domain.User, error) {
	var user domain.User
	query := "SELECT id, username, password FROM users WHERE username = $1"
	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password) // тут косякнул (query, username) username это входные данные
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil // тут забыл сам написать успешную отработку но пофиксил(значит что то нашли по логике далее это ошибка)
}

func (r *PostgresUserRepo) Create(user *domain.User) (*domain.User, error) {
	query := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id"
	err := r.db.QueryRow(query, user.Username, user.Password).Scan(&user.ID) // ничего не возвращает только пишет данные в структуру, ну и вернет nil если все ок так как возврщает только error
	if err != nil {
		return nil, err
	}
	return user, nil
}
