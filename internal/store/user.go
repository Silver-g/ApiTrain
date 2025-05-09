package store

import (
	"ApiTrain/internal/domain"
)

type UserRepository interface {
	Create(user *domain.User) (*domain.User, error)
}
