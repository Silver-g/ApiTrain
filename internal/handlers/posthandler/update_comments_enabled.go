package posthandler

import (
	"ApiTrain/internal/boundary"
	"ApiTrain/internal/security"
	"ApiTrain/internal/service/postservice"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type HandlerUpdateCommentsEnabled struct {
	UpdateCommentsEnabledService postservice.PostService
}

func NewUpdateCommentsEnabledHandler(ucesvc postservice.PostService) *HandlerUpdateCommentsEnabled {
	var newHandlerUpdateData HandlerUpdateCommentsEnabled
	newHandlerUpdateData.UpdateCommentsEnabledService = ucesvc
	return &newHandlerUpdateData
}

func (h *HandlerUpdateCommentsEnabled) UpdateCommentsEnabled(w http.ResponseWriter, r *http.Request) {
	var err error
	decoder := json.NewDecoder(r.Body)
	if r.Method != http.MethodPatch {
		boundary.WriteResponseErr(w, 405, boundary.ErrorResponse{
			ErrorCode: "MethodNotAllowed",
			Message:   "Only PATCH method is allowed.",
		})
		return
	}
	//
	authHeader := r.Header.Get("Authorization")
	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenStr == authHeader {
		boundary.WriteResponseErr(w, 401, boundary.ErrorResponse{
			ErrorCode: "Unauthorized",
			Message:   "Invalid Authorization header format",
		})
		return
	}
	userId, err := security.ParseJwt(tokenStr)
	//
	if err != nil {
		boundary.WriteResponseErr(w, 401, boundary.ErrorResponse{
			ErrorCode: "Unauthorized",
			Message:   "Invalid token",
		})
		return
	}
	var postIdStr string
	parts := strings.Split(r.URL.Path, "/") // ["", "posts", "123", "comments"]
	if len(parts) < 4 || parts[3] != "comments-enabled" {
		boundary.WriteResponseErr(w, 400, boundary.ErrorResponse{
			ErrorCode: "StatusBadRequest",
			Message:   "invalidRequest",
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

	var enabledDataReq boundary.UpdatePostRequest //почему new уточнить
	err = decoder.Decode(&enabledDataReq)
	if err != nil {
		boundary.WriteResponseErr(w, 400, boundary.ErrorResponse{
			ErrorCode: "StatusBadRequest",
			Message:   "Invalid syntax",
		})
		return
	}
	err = boundary.UpdateCommentsEnabledValidate(enabledDataReq)
	if err != nil {
		boundary.WriteResponseErr(w, 400, boundary.ErrorResponse{
			ErrorCode: "ValidationError",
			Message:   err.Error(),
		})
		return
	}
	enabledDataMaping := boundary.UpdateCommentsEnabledMaping(&enabledDataReq, postId, userId)
	updatedData, err := h.UpdateCommentsEnabledService.UpdatePostCommentsEnabled(&enabledDataMaping)
	if err != nil {
		boundary.WriteResponseErr(w, 500, boundary.ErrorResponse{
			ErrorCode: "InternalError",
			Message:   err.Error(), //ленивая обработка тут 2 сценария доделать
		})
		return
	}
	boundary.WriteResponseSuccess(w, 200, boundary.SuccessResponse{
		ResponseData: updatedData,
		Message:      "Успешное обновление",
	})
}
