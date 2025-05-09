package handlers

import (
	"ApiTrain/internal/boundary"
	"ApiTrain/internal/dto"
	"ApiTrain/internal/service/userService"
	"encoding/json"
	"fmt"
	"net/http"
)

//	type Handler struct {
//		DB *sql.DB
//	}
type Handler struct {
	Service userService.UserRegister
}

func NewHandler(svc userService.UserRegister) *Handler {
	var newHandlerex Handler
	newHandlerex.Service = svc
	return &newHandlerex
}
func WriteResponseSuccess(w http.ResponseWriter, statusCode int, successResp dto.SuccessResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(successResp)
}

func WriteResponseErr(w http.ResponseWriter, statusCode int, errResp dto.ErrorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errResp)
}
func (h *Handler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) { // уточнить почему тут делаем ссылку  r *http.Request
	decoder := json.NewDecoder(r.Body) // константы????
	//var errResp ErrorResponse пока отключил так как функция есть
	if r.Method != http.MethodPost { // r.Metgod - то что прислал клиент в запросе, http.MetgodPost - константа которая хранит метод пост
		WriteResponseErr(w, 405, dto.ErrorResponse{
			ErrorCode: "MethodNotAllowed",
			Message:   "Only POST method is allowed.",
		})
		return
	}
	var userReq boundary.UserRequest
	err := decoder.Decode(&userReq) // уточнить почему тут делаем ссылку &user
	if err != nil {
		WriteResponseErr(w, 400, dto.ErrorResponse{
			ErrorCode: "StatusBadRequest",
			Message:   "Не верный синтаксис",
		})
		return
	}

	err = boundary.UserValidate(userReq)
	if err != nil {
		WriteResponseErr(w, 400, dto.ErrorResponse{
			ErrorCode: "ValidationError",
			Message:   err.Error(),
		})
		return
	}
	userMaping := boundary.ConvertToInternal(userReq)
	_, err = h.Service.Register(userMaping)
	if err != nil {
		fmt.Println("ошибка при создании пользователя:", err) // добавь лог
		WriteResponseErr(w, 500, dto.ErrorResponse{
			ErrorCode: "InternalError",
			Message:   "Failed to create user",
		})
		return
	}

	WriteResponseSuccess(w, 200, dto.SuccessResponse{
		Message: "Пользователь успешно создан",
	})
}

//Конспект ошибки +
//КАК ПРАВИЛЬНО ОБРАБАТЫВАТЬ ОШИБКИ УЗНАТЬ
// тоби надо написать функцию которая в качестве аргументов принимает 3 параметра статус код сообщения для ошибок можно в виде констант
// так же нужно подумать над запиханием туда заголовка пока что запиши его без доп параметров чтобы ну не какать потом пишешь ну и пишешь еще в конце
// функции encoder.Encode(errResp) пока все думай лох

// encoder := json.NewEncoder(w)
// if r.Method != http.MethodPost {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(405)
// 	errResp := ErrorResponse{
// 		ErrorInfo: "MetodIsNotAllowed",
// 		Message:   "Only POST method is allowed.",
// 	}
// 	encoder.Encode(errResp)
// }

//ниже валидация
// if user.UserName == "" || user.Password == "" {
// 	writeResponseErr(w, 400, ErrorResponse{
// 		ErrorCode: "ValidationError",
// 		Message:   "Login and password cannot be empty",
// 	})
// 	return
// }
// if len(user.UserName) < 5 || len(user.UserName) > 25 {
// 	writeResponseErr(w, 400, ErrorResponse{
// 		ErrorCode: "ValidationError",
// 		Message:   "Username must be between 5 and 25 characters.",
// 	})
// 	return
// }
// if len(user.Password) > 15 {
// 	writeResponseErr(w, 400, ErrorResponse{
// 		ErrorCode: "ValidationError",
// 		Message:   "Password must be at least 15 characters.",
// 	})
// 	return
// }
//НЕОБХОДИМО ДОПИСАТЬ ЛОГИКУ ВАЛИДАЦИИ НО ТОЛЬКО ПОСЛЕ ТОГО КАК РАЗБЕРЕШЬСЯ С БД и писать валидацию будешь в отдельную функцию или пакеты пока не понятно
// тут логика отправки в бд
