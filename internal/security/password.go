package security

import "golang.org/x/crypto/bcrypt"

//для регистрации
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

//для логина
func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// мб вообще делать не функцию крейт то то а сделать это через метод интерфейса пока сложно сказать что мне это даст
