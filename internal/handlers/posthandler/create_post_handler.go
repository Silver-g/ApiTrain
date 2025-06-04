package posthandler

import (
	"ApiTrain/internal/boundary"
	"ApiTrain/internal/security"
	"ApiTrain/internal/service/postservice"
	"encoding/json"
	"net/http"
	"strings"
)

type HandlerCreatePost struct {
	CreatePostService postservice.CreatePost
}

func NewCreatePostHandler(cpsvc postservice.CreatePost) *HandlerCreatePost {
	var newHandlerex HandlerCreatePost
	newHandlerex.CreatePostService = cpsvc
	return &newHandlerex
}
func (h *HandlerCreatePost) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var err error
	if r.Method != http.MethodPost {
		boundary.WriteResponseErr(w, 405, boundary.ErrorResponse{
			ErrorCode: "MethodNotAllowed",
			Message:   "Only POST method is allowed.",
		})
		return
	}
	authHeader := r.Header.Get("Authorization") //извлекаем токен из заголовка мб тут нужно подшаманить чуть

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ") //чистим извелеченные данные удаляем лишнее
	if tokenStr == authHeader {
		boundary.WriteResponseErr(w, 401, boundary.ErrorResponse{
			ErrorCode: "Unauthorized",
			Message:   "Invalid Authorization header format",
		})
		return
	}
	userId, err := security.ParseJwt(tokenStr) //парсим токен
	if err != nil {
		boundary.WriteResponseErr(w, 401, boundary.ErrorResponse{
			ErrorCode: "Unauthorized",
			Message:   "Invalid token",
		})
		return
	}

	postReq := new(boundary.CreatePostRequest) // var postReq *boundary.CreatePostRequest тут конспект + разобратся с указателем почему так там было что то про памят мб переделать выглядит страшна
	err = decoder.Decode(&postReq)             //вспомнить про new вроде бы для указателя на поле comment enabled
	if err != nil {
		boundary.WriteResponseErr(w, 400, boundary.ErrorResponse{
			ErrorCode: "StatusBadRequest",
			Message:   "Invalid syntax",
		})
		return
	}
	err = boundary.CreatePostPostValidate(*postReq) //добавил указатель для валидации bool уточнить перепроверить
	if err != nil {
		boundary.WriteResponseErr(w, 400, boundary.ErrorResponse{
			ErrorCode: "ValidationError",
			Message:   err.Error(),
		})
		return
	}
	CreatePostMaping := boundary.CreapePostMaping(postReq, userId) //мб завези обрбаотку ошибок
	postId, err := h.CreatePostService.PostCreate(&CreatePostMaping)
	if err != nil {
		if err == postservice.ErrAlreadyExist {
			boundary.WriteResponseErr(w, 409, boundary.ErrorResponse{
				ErrorCode: "Conflict",
				Message:   err.Error(),
			})
			return

		}
		boundary.WriteResponseErr(w, 500, boundary.ErrorResponse{
			ErrorCode: "InternalError",
			Message:   err.Error(),
		})
		return
	}
	boundary.WriteResponseSuccess(w, 200, boundary.SuccessResponse{
		ResponseData: postId,
		Message:      "Пост создан успешно гцгцгц",
	})
}
