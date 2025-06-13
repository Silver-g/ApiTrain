package userrepo

import (
	"ApiTrain/internal/domain"
	"database/sql"
	"errors"
)

var ErrUserNotFound = errors.New("user not found") // это потом убрать в отдельное место
type UserPostgres struct {                         //сделал так как у умного дяди но мб вообще нужно иначе ЮЛЯ В ПОМОЩЬ
	db *sql.DB
}

func NewPostgresUser(dataBase *sql.DB) *UserPostgres {
	var userRepoPointer UserPostgres
	userRepoPointer.db = dataBase
	return &userRepoPointer
} // ебучий тильт чтобы оно работало нужно переносить локигку подлючения к базе или это в отдельный файл я в ахуе....

// ну я использовал структуру из другого файла и опять же нужно делать новое соединение новую структуру и новую функцию конструктор
// или вынести это в какой то отедльный файлик и передавать в репозитории как глобальное решение (я хз как надо памагите)
// повторить синтаксис как писать кастомные ошибки
func (r *UserPostgres) LoginByUsername(username string) (*domain.LoginUserInternal, error) {
	var userLogin domain.LoginUserInternal
	query := "SELECT id, username, password FROM users WHERE username = $1" // переписать сеодня в конспект напомнить sql логику не смог дописать услвоие поиска
	err := r.db.QueryRow(query, username).Scan(&userLogin.Id, &userLogin.Username, &userLogin.PasswordHash)
	// тут накосячил фул с синтаксисом повторить весь аргменты и после скан логику тоже что нужно объявлять экземпляр и с ним работать
	// забыл что функция сама по себе ничего не возвращает она просто пишет из этого и следует логика (возвраащет только ошибку если что то пощло не так)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &userLogin, nil
}
func (r *UserPostgres) GetUserById(userid int) error {
	var exists bool
	query := "SELECT EXISTS (SELECT 1 FROM user WHERE id = $1)"
	err := r.db.QueryRow(query, userid).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return ErrUserNotFound
	}
	return nil
}
func (r *UserPostgres) GetByUsername(username string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS (SELECT 1 FROM users WHERE username = $1)"
	err := r.db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		return false, err //тут мб проблемы с логикой
	}
	return exists, nil
}

func (r *UserPostgres) Create(user *domain.User) (*domain.User, error) {
	query := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id"
	err := r.db.QueryRow(query, user.Username, user.Password).Scan(&user.Id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
