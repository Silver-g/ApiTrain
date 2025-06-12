package commenthandler

import (
	"ApiTrain/internal/boundary"
	"ApiTrain/internal/service/commentservice"
	"net/http"
	"strconv"
	"strings"
)

type HandlerGetPostComments struct {
	GetPostCommentsResponse commentservice.CommentService
}

func NewBuildTreeHandler(ccsvc commentservice.CommentService) *HandlerGetPostComments { //провреить необходимость но мб надо из за сложный случай
	var newHandlerBuildTree HandlerGetPostComments
	newHandlerBuildTree.GetPostCommentsResponse = ccsvc
	return &newHandlerBuildTree
}

func (h *HandlerGetPostComments) BuildTreeHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Method != http.MethodGet {
		boundary.WriteResponseErr(w, 405, boundary.ErrorResponse{
			ErrorCode: "MethodNotAllowed",
			Message:   "Only POST method is allowed.",
		})
		return
	}
	////////////////////////////////////////////////
	partsStr := strings.Split(r.URL.Path, "/")

	if len(partsStr) < 4 || partsStr[3] != "comments" { //мб другое условие типо == 4
		boundary.WriteResponseErr(w, 400, boundary.ErrorResponse{
			ErrorCode: "BadRequest",
			Message:   "Invalid path: expected /posts/{id}/comments",
		})
		return
	}
	idStr := partsStr[2]
	postId, err := strconv.Atoi(idStr)
	if err != nil {
		boundary.WriteResponseErr(w, 400, boundary.ErrorResponse{
			ErrorCode: "BadRequest",
			Message:   "Invalid post ID",
		})
		return
	}
	if postId <= 0 {
		boundary.WriteResponseErr(w, 400, boundary.ErrorResponse{
			ErrorCode: "StatusBadRequest",
			Message:   "post_id must be a positive integer",
		})
		return
	} //ручками перепиши блок чтобы не втыкал в конце
	////////////////////////////////////////////////////////
	comentsData, err := h.GetPostCommentsResponse.GetCommentsData(postId)
	if err != nil {
		boundary.WriteResponseErr(w, 400, boundary.ErrorResponse{
			ErrorCode: "BadRequest",
			Message:   err.Error(),
		})
		return
	}
	commentsTree, err := commentservice.CreateTreeComments(comentsData)
	if err != nil {
		boundary.WriteResponseErr(w, 400, boundary.ErrorResponse{
			ErrorCode: "BadRequest",
			Message:   err.Error(),
		})
		return
	}
	postData, err := h.GetPostCommentsResponse.GetPostByIdData(postId)
	if err != nil {
		boundary.WriteResponseErr(w, 400, boundary.ErrorResponse{
			ErrorCode: "BadRequest",
			Message:   err.Error(),
		})
		return //просто жесть надо чистить ближайшее время дальше это не может в свалку превращатся маперы ошибки и тд все нужно прибрать разделить ответственность и тд
	}
	CommentsAndPostTree, err := commentservice.CreateResponsePostAndComments(postData, commentsTree)
	if err != nil {
		boundary.WriteResponseErr(w, 400, boundary.ErrorResponse{
			ErrorCode: "BadRequest",
			Message:   err.Error(),
		})
		return
	} //поправить ошибку
	boundary.WriteResponseSuccess(w, 200, boundary.SuccessResponse{
		ResponseData: CommentsAndPostTree,
		Message:      "Ответ с дервом получен гцгц",
	})
}
