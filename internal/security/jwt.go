package security

import (
	"encoding/base64"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID int) (string, error) {
	var err error
	base64SecretJwt := os.Getenv("JWT_SECRET_BASE64")
	secretKey, err := base64.StdEncoding.DecodeString(base64SecretJwt)
	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := jwtToken.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
