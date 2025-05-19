package postgres

import (
	"ApiTrain/internal/domain"
	"database/sql"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(dataBase *sql.DB) *Postgres {
	var userRepoPointer Postgres
	userRepoPointer.db = dataBase
	return &userRepoPointer
} // ебучий тильт чтобы оно работало нужно переносить локигку подлючения к базе или это в отдельный файл я в ахуе....

// ну я использовал структуру из другого файла и опять же нужно делать новое соединение новую структуру и новую функцию конструктор
// или вынести это в какой то отедльный файлик и передавать в репозитории как глобальное решение (я хз как надо памагите)
// повторить синтаксис как писать кастомные ошибки
func (r *Postgres) LoginByUsername(username string) (*domain.LoginUserInternal, error) {
	var userLogin domain.LoginUserInternal
	query := "SELECT id, username, password FROM users WHERE username = $1" // переписать сеодня в конспект напомнить sql логику не смог дописать услвоие поиска
	err := r.db.QueryRow(query, username).Scan(&userLogin.ID, &userLogin.Username, &userLogin.PasswordHash)
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
