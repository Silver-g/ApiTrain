package boundary

import (
	"ApiTrain/internal/domain"
	"errors"
)

var prohibitedChars = map[rune]struct{}{
	'\'': {}, '"': {}, ';': {}, '\\': {}, '`': {}, '=': {},
	'(': {}, ')': {}, '{': {}, '}': {}, '[': {}, ']': {},
}

func isOnlySpacesOrUnderscores(s string) bool {
	for _, ch := range s {
		if ch != ' ' && ch != '_' && ch != '\t' {
			return false
		}
	}
	return true
}

func containsProhibitedChars(s string) bool {
	for _, ch := range s {
		if _, exists := prohibitedChars[ch]; exists {
			return true
		}
	}
	return false
}

// внешняя структура
type UserLogin struct {
	Username string `json:"username"`
	Password string `json: "password"`
}

// Maping
func UserLoginMaping(userLoginReq UserLogin) domain.LoginUserInternal {
	return domain.LoginUserInternal{
		Username: userLoginReq.Username,
		Password: userLoginReq.Password,
	}
}

// Validate
func LoginValidate(userLoginReq UserLogin) error {
	if userLoginReq.Username == "" || userLoginReq.Password == "" {
		return errors.New("Login and password cannot be empty")
	}
	if isOnlySpacesOrUnderscores(userLoginReq.Username) || isOnlySpacesOrUnderscores(userLoginReq.Password) {
		return errors.New("login and password cannot consist only of spaces or underscores")
	}

	if containsProhibitedChars(userLoginReq.Username) || containsProhibitedChars(userLoginReq.Password) {
		return errors.New("login or password contains prohibited characters")
	} //почистить вывод ошибок
	return nil
}
