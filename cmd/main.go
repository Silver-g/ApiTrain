package main

import (
	"ApiTrain/internal/handlers"
	"ApiTrain/internal/service/postService"
	"ApiTrain/internal/service/userService"
	"ApiTrain/internal/store"
	"ApiTrain/internal/store/postgres"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func ServerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "йоу")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}
	// Подключаемся к БД
	db, err := store.ConnectDB()
	if err != nil {
		panic(err)
	}
	// Создаём репозиторий и сервис
	repo := postgres.NewPostgres(db)
	svc := userService.NewUserService(repo)
	cpsvc := postService.CreatePostServiceConstruct(repo)
	// Создаём HTTP-обработчик
	handler := handlers.NewHandler(svc)
	handlerlogin := handlers.LoginHandler(svc)
	handlerPostCreate := handlers.CreatePostHandler(cpsvc)
	http.HandleFunc("/", ServerHandler)
	http.HandleFunc("/register", handler.RegisterUserHandler)
	http.HandleFunc("/login", handlerlogin.LoginUserHandler)
	http.HandleFunc("/createpost", handlerPostCreate.CreatePostHandler) // с именами пиздец
	fmt.Println("Server running on http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ошибка при запуске")
	}
}
