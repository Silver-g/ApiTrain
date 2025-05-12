package userService

import (
	"ApiTrain/internal/domain"
	"ApiTrain/internal/security"
	"ApiTrain/internal/store"
	"errors"
)

var ErrUserAlreadyExists = errors.New("user already exists") // сделать константой и пернести в отдельный файл как и все ошибки

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
	existingUser, err := s.userRepo.GetByUsername(user.Username)
	if err != nil {
		// Внутренняя ошибка при попытке получить пользователя
		return nil, err
	}
	if existingUser != nil {
		// Возвращаем конкретную ошибку если пользователь уже существует
		return nil, ErrUserAlreadyExists
	}
	hashedPassword, err := security.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword
	createdUser, err := s.userRepo.Create(&user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}
