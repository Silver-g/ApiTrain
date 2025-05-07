package store

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Ошибка загрузки .env файла")
		return nil, err
	}
	connectString := "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"
	connectInput := fmt.Sprintf(
		connectString,
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	driver := os.Getenv("DB_DRIVER")
	db, err := sql.Open(driver, connectInput)
	if err != nil {
		fmt.Println("ошибка подключения")
		return nil, fmt.Errorf("не удалось открыть соединение: %w", err)
	}

	return db, nil
}
