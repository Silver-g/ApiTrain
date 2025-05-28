package main

import (
	"ApiTrain/internal/handlers"
	"ApiTrain/internal/service/commentService"
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
	cpsvc := postService.NewCreatePostService(repo)
	ccsvc := commentService.NewCreateCommentService(repo)
	// Создаём HTTP-обработчик
	handler := handlers.NewHandler(svc)
	handlerlogin := handlers.LoginHandler(svc)
	handlerPostCreate := handlers.CreatePostHandler(cpsvc)
	handlerCommentCreate := handlers.CreateCommentHandler(ccsvc)
	http.HandleFunc("/", ServerHandler)
	http.HandleFunc("/register", handler.RegisterUserHandler)
	http.HandleFunc("/login", handlerlogin.LoginUserHandler)
	http.HandleFunc("/createpost", handlerPostCreate.CreatePostHandler)          // с именами пиздец
	http.HandleFunc("/createcomment", handlerCommentCreate.CreateCommentHandler) // с именами полная жесть как она есть намекну проблема в схожих именах назввания метода обработчика функции конструктора и тд тут нужен метод обработчика
	fmt.Println("Server running on http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ошибка при запуске")
	}
}
