package userhandler

import (
	"ApiTrain/internal/boundary"
	"ApiTrain/internal/service/userservice"
	"encoding/json"
	"net/http"
)

// Так либовски тут был замечен косяк а точне ты описываешь структру реализуешь обработчик через метод
// но не описываешь интерфейс который реализует этот метод а все потому что ты кастылем херачишь
// роут в main следовательно что нужно узнать у Юли как тут правильно поступить по логике нужно описать интрфейс
// и куда то не понятно куда вынести его в отедельный файл, в нем собрать метод регистрации лоигна и других твоих обработчиков
// где это хранить да шут его знает либо в сервисе либо в пакете обработчика, костыль в main это (http.HandleFunc) а нужно в теории (mux.HandleFunc)
type HandlerUserRegister struct {
	Service userservice.UserRegister
}

func NewHandlerRegister(svc userservice.UserRegister) *HandlerUserRegister {
	var newHandUser HandlerUserRegister
	newHandUser.Service = svc
	return &newHandUser
}

func (h *HandlerUserRegister) RegisterUserHandler(w http.ResponseWriter, r *http.Request) { // уточнить почему тут делаем ссылку  r *http.Request
	decoder := json.NewDecoder(r.Body) // константы????
	//var errResp ErrorResponse пока отключил так как функция есть
	if r.Method != http.MethodPost { // r.Metgod - то что прислал клиент в запросе, http.MetgodPost - константа которая хранит метод пост
		boundary.WriteResponseErr(w, 405, boundary.ErrorResponse{
			ErrorCode: "MethodNotAllowed",
			Message:   "Only POST method is allowed.",
		})
		return
	}
	var userReq boundary.UserRequest
	err := decoder.Decode(&userReq) // уточнить почему тут делаем ссылку &user
	if err != nil {
		boundary.WriteResponseErr(w, 400, boundary.ErrorResponse{
			ErrorCode: "StatusBadRequest",
			Message:   "Не верный синтаксис",
		})
		return
	}

	err = boundary.UserValidate(userReq)
	if err != nil {
		boundary.WriteResponseErr(w, 400, boundary.ErrorResponse{
			ErrorCode: "ValidationError",
			Message:   err.Error(),
		})
		return
	}
	userMaping := boundary.ConvertToInternal(userReq)
	_, err = h.Service.Register(userMaping) // тут игонрируем 1ое поле так как мы не возвращаем пользователю данных о созадном пользователе ни username ни id и тд
	if err != nil {
		// Обрабатываем ошибку, если пользователь уже существует
		if err == userservice.ErrUserAlreadyExists {
			boundary.WriteResponseErr(w, 400, boundary.ErrorResponse{
				ErrorCode: "UserAlreadyExists",
				Message:   "User with this username already exists.",
			})
			return
		}
		// Обрабатываем другие ошибки
		boundary.WriteResponseErr(w, 500, boundary.ErrorResponse{
			ErrorCode: "InternalError",
			Message:   "Failed to create user",
		})
		return
	}
	boundary.WriteResponseSuccess(w, 200, boundary.SuccessResponse{
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
