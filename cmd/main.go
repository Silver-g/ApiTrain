package main

import (
	"ApiTrain/internal/config"
	"ApiTrain/internal/handlers/commenthandler"
	"ApiTrain/internal/handlers/posthandler"
	"ApiTrain/internal/handlers/userhandler"
	"ApiTrain/internal/service/commentservice"
	"ApiTrain/internal/service/postservice"
	"ApiTrain/internal/service/userservice"
	"ApiTrain/internal/store/db"
	"ApiTrain/internal/store/postgres/commentrepo"
	"ApiTrain/internal/store/postgres/postrepo"
	"ApiTrain/internal/store/postgres/userrepo"
	"fmt"
	"log"
	"net/http"
)

func ServerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "йоу")
}

func main() {
	var err error
	//Подгурзка енв украл из db(сделал отдельную функцию, это логично так как может быть несколько енв файлов)
	err = config.InitConfig(".env") // имя передал напрямую по факту норм но вообще можно вынести его в константу
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла") // лог фатал завершит программу через (os.Exit(1)) эта штука принудительно завершит программу. Исппользую именно там где некуда уже передавать ошибу нет четкого механизма ответа нужно просто сказать что что то критично пошло не так паника или типо того.
	}
	// Подключаемся к БД
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Не удалось открыть соединение с базой данных")
	}
	// Создаём репозиторий и сервис
	userRepo := userrepo.NewPostgresUser(db)
	postRepo := postrepo.NewPostgresPost(db)
	commentRepo := commentrepo.NewPostgresComment(db)
	////////////////////////////////////////////////////////
	svc := userservice.NewUserService(userRepo)
	cpsvc := postservice.NewPostService(postRepo)
	ccsvc := commentservice.NewCommentService(commentRepo)
	// Создаём HTTP-обработчик
	handler := userhandler.NewHandlerRegister(svc)
	handlerlogin := userhandler.NewLoginHandler(svc)
	//
	handlerPostCreate := posthandler.NewCreatePostHandler(cpsvc)
	//
	handlerCommentCreate := commenthandler.NewCreateCommentHandler(ccsvc)
	handlerBuildTree := commenthandler.NewBuildTreeHandler(ccsvc)
	//
	var temp commenthandler.CommentRouter
	temp.CreateHandler = handlerCommentCreate
	temp.TreeHandler = handlerBuildTree //разобрать эту запись
	commentRouter := &temp
	//////////////////////////////////////////////////////////
	http.HandleFunc("/", ServerHandler)
	http.HandleFunc("/register", handler.RegisterUserHandler)
	http.HandleFunc("/login", handlerlogin.LoginUserHandler) // тут кста все же не на функцию конструктор а как ты и думал на метод обработчика что в целомл логично
	http.HandleFunc("/posts", handlerPostCreate.CreatePostHandler)
	// ниже костыль
	http.Handle("/posts/", commentRouter) //Подробно разобрать что как и почему в самом кастомном роуте и в целом
	fmt.Println("Server running on http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ошибка при запуске")
	}
}
