package userService

import (
	"ApiTrain/internal/domain"
	"ApiTrain/internal/security"
	"ApiTrain/internal/store"
	"ApiTrain/internal/store/postgres"
	"errors"
)

var ErrUserAlreadyExists = errors.New("user already exists") //  пернести в отдельный файл как и все ошибки

type UserRegister interface {
	Register(user domain.User) (*domain.User, error)
}
type UserService struct {
	userRepo store.UserRepository
}

func NewUserService(repo store.UserRepository) *UserService {
	var userPointerBd UserService
	userPointerBd.userRepo = repo
	return &userPointerBd
}

func (s *UserService) Register(user domain.User) (*domain.User, error) {
	existingUser, err := s.userRepo.GetByUsername(user.Username)
	if err != nil && err != postgres.ErrUserNotFound {
		// Внутренняя ошибка при попытке получить пользователя
		return nil, err
	} // мб переделать обработку ошибок с error.is
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
