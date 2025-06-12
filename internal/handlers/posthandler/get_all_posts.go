package posthandler

import (
	"ApiTrain/internal/boundary"
	"ApiTrain/internal/service/postservice"
	"net/http"
)

type HandlerGetAllPosts struct {
	AllPostsGetService postservice.PostService
}

func NewGetPostHandler(cpsvc postservice.PostService) *HandlerGetAllPosts {
	var newHandlerex HandlerGetAllPosts
	newHandlerex.AllPostsGetService = cpsvc //мб можно объеденеить логику с общей структурой там где создаавал пост
	return &newHandlerex
}
func (h *HandlerGetAllPosts) GetAllPostsHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Method != http.MethodGet {
		boundary.WriteResponseErr(w, 405, boundary.ErrorResponse{
			ErrorCode: "MethodNotAllowed",
			Message:   "Only GET method is allowed.",
		})
		return
	}
	// partsStr := strings.Split(r.URL.Path, "/")
	// if len(partsStr) > 2 && partsStr[1] != "posts" { // я думаю это избыточно сильно избыточно
	// 	boundary.WriteResponseErr(w, 400, boundary.ErrorResponse{
	// 		ErrorCode: "BadRequest",
	// 		Message:   "Invalid path: expected /posts",
	// 	})
	// 	return
	// }
	postsList, err := h.AllPostsGetService.GetAllPostsServ()
	if err != nil {
		boundary.WriteResponseErr(w, 500, boundary.ErrorResponse{
			ErrorCode: "InternalError",
			Message:   err.Error(),
		})
		return
	}
	boundary.WriteResponseSuccess(w, 200, boundary.SuccessResponse{
		ResponseData: postsList,
		Message:      "Список получен успешно гцгцгц",
	})
}
