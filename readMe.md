## ApiTrain

Учебный проект на Go с авторизацией, постами и комментариями.

### Структура проекта

- `internal/domain` — доменные сущности.
- `internal/store` — слой хранения (PostgreSQL).
- `internal/service` — бизнес-логика.
- `cmd/main.go` — точка входа.

### Запуск проекта
go run cmd/main.go


### Основные запросы:
Ниже представленные основные запросы, тестировал через insomnia
#### register регистрация
POST http://localhost:8080/register
{
    "login": "Adminuser", 
    "password": "qazwsxedcrfvtgbyhn"  
}
#### login логин
POST http://localhost:8080/login
{
    "username": "Adminuser", 
    "password": "qazwsxedcrfvtgbyhn"  
}
Для токена: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDkxMTAzMDcsInVzZXJfaWQiOjIxfQ.5q3CaQqHhblnapTXGRZEuqt00OQUV19VImutMnfmUfo
#### Создание поста /posts 
POST http://localhost:8080/posts
{
    "title": "Тестовый заголовок", 
    "text": "Тестовый текст поста 55",
    "comments_enabled": true
}
#### Создание комментария 
POST http://localhost:8080/posts/13/comments
{
    "text": "Тестовый коммент к 13 посту", 
    "parentId": null
}
#### Получение списка комментариев и пост
GET http://localhost:8080/posts/13/comments
{
    "text": "Тестовый коммент к 13 посту + ответ + ответ на ответ на ответ ", 
    "parentId": 21
}