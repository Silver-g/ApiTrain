package security

import (
	"encoding/base64"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var UnexpectedMethod error = errors.New("unexpected signing method")
var FailedParseClaims error = errors.New("failed to parse claims")
var InvalidJwtToken error = errors.New("invalid token")
var UnreadableUserId error = errors.New("user_id is missing or not a number")

func ParseJwt(jwtToken string) (int, error) { // удачи сегодня вечером разобрать код этого нечто
	var err error
	base64SecretJwt := os.Getenv("JWT_SECRET_BASE64")
	secretKey, err := base64.StdEncoding.DecodeString(base64SecretJwt)
	if err != nil {
		return 0, err
	}
	token, err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
		method := t.Method
		_, ok := method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, UnexpectedMethod
		}
		return secretKey, nil
	})
	if err != nil {
		return 0, err
	}
	claimsRaw := token.Claims
	claims, ok := claimsRaw.(jwt.MapClaims)
	if !ok {
		return 0, FailedParseClaims
	}

	if !token.Valid {
		return 0, InvalidJwtToken
	}

	userIDFloat, ok := claims["user_id"].(float64) //нахуя можно же просто инт айдишник не будет дробью никогда
	if !ok {
		return 0, UnreadableUserId
	}

	return int(userIDFloat), nil
}

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
