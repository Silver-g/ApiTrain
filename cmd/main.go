package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) { // w http.ResponseWriter это буквально ответ на зпрос клиента
	fmt.Fprintln(w, "Привет, сервер на Go работает!") //r *http.Request это буквально сам запрос
} //fmt.Fprintln(w, "Привет, сервер на Go работает!")  буквально пишем ответ на запрос /

func main() {
	http.HandleFunc("/", helloHandler) // эта функция описывает роуты и ссылатется на выше описанный метод который выполняется на этом роуте

	fmt.Println("Сервер запущен на http://localhost:8080")
	err := http.ListenAndServe(":8080", nil) // Это главная функция, которая запускает HTTP-сервер.
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
