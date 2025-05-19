package handlers

import (
	"ApiTrain/internal/boundary"
	"ApiTrain/internal/service/userService"
	"encoding/json"
	"net/http"
)

type HandlerLogin struct {
	ServiceLogin userService.UserLogin
}

// Так либовски тут был замечен косяк а точне ты описываешь структру реализуешь обработчик через метод
// но не описываешь интерфейс который реализует этот метод а все потому что ты кастылем херачишь
// роут в main следовательно что нужно узнать у Юли как тут правильно поступить по логике нужно описать интрфейс
// и куда то не понятно куда вынести его в отедельный файл, в нем собрать метод регистрации лоигна и других твоих обработчиков
// где это хранить да шут его знает либо в сервисе либо в пакете обработчика, костыль в main это (http.HandleFunc) а нужно в теории (mux.HandleFunc)
func LoginHandler(svc userService.UserLogin) *HandlerLogin {
	var newHandlerex HandlerLogin
	newHandlerex.ServiceLogin = svc
	return &newHandlerex
}
func (h *HandlerLogin) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var err error
	if r.Method != http.MethodPost {
		boundary.WriteResponseErr(w, 405, boundary.ErrorResponse{
			ErrorCode: "MethodNotAllowed",
			Message:   "Only POST method is allowed.",
		})
		return
	}
	var userLoginReq boundary.UserLogin
	err = decoder.Decode(&userLoginReq)
	if err != nil {
		boundary.WriteResponseErr(w, 400, boundary.ErrorResponse{
			ErrorCode: "StatusBadRequest",
			Message:   "Не верный синтаксис",
		})
		return
	}
	err = boundary.LoginValidate(userLoginReq)
	if err != nil {
		boundary.WriteResponseErr(w, 400, boundary.ErrorResponse{
			ErrorCode: "ValidationError",
			Message:   err.Error(),
		})
		return
	}
	loginUserMaping := boundary.UserLoginMaping(userLoginReq)

	userAuthToken, err := h.ServiceLogin.Login(loginUserMaping)

	if err != nil {
		if err == userService.InvalidUsername {
			boundary.WriteResponseErr(w, 404, boundary.ErrorResponse{
				ErrorCode: "AuthError",
				Message:   err.Error(),
			})
			return
		}
		if err == userService.IncorrectPassword {
			boundary.WriteResponseErr(w, 404, boundary.ErrorResponse{
				ErrorCode: "AuthError",
				Message:   err.Error(),
			})
			return
		}
		boundary.WriteResponseErr(w, 500, boundary.ErrorResponse{
			ErrorCode: "AuthError",
			Message:   err.Error(),
		})
		return
	}
	boundary.WriteResponseSuccess(w, 200, boundary.SuccessResponse{
		Token:   userAuthToken,
		Message: "Авторизация прошла успешно",
	})
}
