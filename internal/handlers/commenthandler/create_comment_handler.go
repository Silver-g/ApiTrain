package commenthandler

import (
	"ApiTrain/internal/boundary"
	"ApiTrain/internal/security"
	"ApiTrain/internal/service/commentservice"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type HandlerCreateComment struct {
	CreateCommentService commentservice.CreateCommentServ // Тут передаем интефрейс еще раз напомни себе почему
}

func NewCreateCommentHandler(ccsvc commentservice.CreateCommentServ) *HandlerCreateComment {
	var newHandlerCrCom HandlerCreateComment
	newHandlerCrCom.CreateCommentService = ccsvc
	return &newHandlerCrCom
}
func (h *HandlerCreateComment) CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	decoder := json.NewDecoder(r.Body) //когда ты уже запомнишь что он читает боди сучка
	if r.Method != http.MethodPost {
		boundary.WriteResponseErr(w, 405, boundary.ErrorResponse{
			ErrorCode: "MethodNotAllowed",
			Message:   "Only POST method is allowed.",
		})
		return
	}
	////////////////////////////////////////////////////////////////////////////////////////////////////////
	authHeader := r.Header.Get("Authorization")           // r.Header.Get извелчение заголовка это разобрать как что и как иначе крч сам все знаешь нубчик
	tokenStr := strings.TrimPrefix(authHeader, "Bearer ") // все верно только порядок перепутал повторить работу с strings убейте меня...
	if tokenStr == authHeader {
		boundary.WriteResponseErr(w, 401, boundary.ErrorResponse{ //это надо чтобы проверить что Bearer точно был в строке так надо по стандарту запросов чтобы +- понять что присылают как понял не более чем обще принятая схема или типо того
			ErrorCode: "Unauthorized",
			Message:   "Invalid Authorization header format",
		})
		return
	}
	userId, err := security.ParseJwt(tokenStr)
	if err != nil {
		boundary.WriteResponseErr(w, 401, boundary.ErrorResponse{
			ErrorCode: "Unauthorized",
			Message:   "Invalid token",
		})
		return
	}
	//////////////////////////////////////////////////////////////////////////////////////////////////////// -- РАЗОБРАТЬ идею то я понял но вот синтаксис местами да
	var postIdStr string
	parts := strings.Split(r.URL.Path, "/") // ["", "posts", "123", "comments"]
	if len(parts) < 4 || parts[3] != "comments" {
		boundary.WriteResponseErr(w, 400, boundary.ErrorResponse{
			ErrorCode: "StatusBadRequest", //тут переделал мб вернуть назад
			Message:   "Invalid syntax",
		})
		return
	}
	postIdStr = parts[2]
	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		boundary.WriteResponseErr(w, 400, boundary.ErrorResponse{
			ErrorCode: "StatusBadRequest",
			Message:   "invalid post_id",
		})
		return
	}
	if postId <= 0 {
		boundary.WriteResponseErr(w, 400, boundary.ErrorResponse{
			ErrorCode: "StatusBadRequest",
			Message:   "post_id must be a positive integer",
		})
		return
	}

	//
	var commentReq boundary.CreateCommentRequest
	err = decoder.Decode(&commentReq) // опять забыл указатель курва ты лысая
	if err != nil {
		boundary.WriteResponseErr(w, 400, boundary.ErrorResponse{
			ErrorCode: "StatusBadRequest",
			Message:   "Invalid syntax",
		})
		return
	}
	err = boundary.CreateCommentValidate(commentReq) //почему тут нет указателя я хз//мб нужна отедельно валидация по полю узер айди хотя я думаю что нет так как проверяем уже при парсинге
	if err != nil {
		boundary.WriteResponseErr(w, 400, boundary.ErrorResponse{
			ErrorCode: "ValidationError",
			Message:   err.Error(),
		})
		return
	}
	commentDataMaping := boundary.CreateCommentMaping(commentReq, userId, postId)
	commentId, err := h.CreateCommentService.CommentCreate(&commentDataMaping) //очко с указателями
	if err != nil {
		if err == commentservice.ErrCommentsDisabled {
			boundary.WriteResponseErr(w, 403, boundary.ErrorResponse{
				ErrorCode: "Forbidden",
				Message:   err.Error(), //мб кастомная ошибка нужна, на этапе чистки посмотришь
			})
			return
		}
		if err == commentservice.ErrPostNotFound {
			boundary.WriteResponseErr(w, 404, boundary.ErrorResponse{
				ErrorCode: "NotFound",
				Message:   err.Error(), //мб кастомная ошибка нужна, на этапе чистки посмотришь
			})
			return
		}
		boundary.WriteResponseErr(w, 500, boundary.ErrorResponse{
			ErrorCode: "InternalError",
			Message:   err.Error(), //мб кастомная ошибка нужна, на этапе чистки посмотришь
		})
		return
	}
	boundary.WriteResponseSuccess(w, 201, boundary.SuccessResponse{
		ResponseData: commentId,
		Message:      "Комент создан успешно гцгцгц",
	})
}
