package userservice

import (
	"ApiTrain/internal/domain"
	"ApiTrain/internal/security"
	"ApiTrain/internal/store/repository"
	"errors"
)

var ErrUserAlreadyExists = errors.New("user already exists") //  пернести в отдельный файл как и все ошибки

type UserRegister interface {
	Register(user domain.User) (*domain.User, error)
}
type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	var userPointerBd UserService
	userPointerBd.userRepo = repo
	return &userPointerBd
}

func (s *UserService) Register(user domain.User) (*domain.User, error) {
	existingUser, err := s.userRepo.GetByUsername(user.Username)
	if err != nil {
		return nil, err
	}
	if existingUser == true {
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
