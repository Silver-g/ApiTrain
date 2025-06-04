package userrepo

import (
	"ApiTrain/internal/domain"
)

type UserRepository interface {
	Create(user *domain.User) (*domain.User, error)
	GetByUsername(username string) (bool, error)
	LoginByUsername(username string) (*domain.LoginUserInternal, error)
}
