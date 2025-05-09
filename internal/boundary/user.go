package boundary

import (
	"ApiTrain/internal/domain"
	"errors"
)

type UserRequest struct {
	Username string `json:"login"`
	Password string `json:"password"`
}

// maping
func ConvertToInternal(userRequest UserRequest) domain.User {
	return domain.User{
		Username: userRequest.Username,
		Password: userRequest.Password,
	}
}

// validate
func UserValidate(userReq UserRequest) error {
	if userReq.Username == "" || userReq.Password == "" {
		return errors.New("Login and password cannot be empty")
	}
	if len(userReq.Username) < 5 || len(userReq.Username) > 25 {
		return errors.New("Username must be between 5 and 25 characters")
	}
	if len(userReq.Password) < 15 {
		return errors.New("Password must not exceed 15 characters")
	}
	return nil
}
