package config

import "github.com/joho/godotenv"

func InitConfig(file string) error { // видимпо потому что Load сам принимает строку в себя
	err := godotenv.Load(file) //видимо просто возвращает ничего либо ошибку поэтому логика такаая
	return err
}
