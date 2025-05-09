package userService

import (
	"ApiTrain/internal/domain"
	"ApiTrain/internal/store"
)

type UserRegister interface {
	Register(user domain.User) (*domain.User, error)
}
type userService struct {
	userRepo store.UserRepository
}

func NewUserService(repo store.UserRepository) *userService {
	var userPointerBd userService
	userPointerBd.userRepo = repo
	return &userPointerBd
}

func (s *userService) Register(user domain.User) (*domain.User, error) {
	createdUser, err := s.userRepo.Create(&user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}
