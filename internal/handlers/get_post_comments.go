package handlers

import (
	"ApiTrain/internal/boundary"
	"ApiTrain/internal/service/commentService"
	"net/http"
	"strconv"
)

type HandlerGetPostComments struct {
	GetPostCommentsResponse commentService.CommentService
}

func BuildTreeHandler(ccsvc commentService.CommentService) *HandlerGetPostComments {
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
	idStr := r.URL.Path[len("/posts/"):] // разобарть синтаксис подучить подобные конструкции и запомнить смыл в том что это взятие под строки (ты забыл про такое)
	if idStr == "" {
		boundary.WriteResponseErr(w, 400, boundary.ErrorResponse{
			ErrorCode: "BadRequest",
			Message:   "Post ID is missing",
		})
		return
	}
	postId, err := strconv.Atoi(idStr)
	if err != nil {
		boundary.WriteResponseErr(w, 400, boundary.ErrorResponse{
			ErrorCode: "BadRequest",
			Message:   "Invalid post ID",
		})
	}
	comentsData, err := h.GetPostCommentsResponse.GetCommentsData(postId)
	if err != nil {
		boundary.WriteResponseErr(w, 400, boundary.ErrorResponse{
			ErrorCode: "BadRequest",
			Message:   err.Error(),
		})
		return
	}
	commentsTree, err := commentService.CreateTreeComments(comentsData)
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
	CommentsAndPostTree, err := commentService.CreateResponsePostAndComments(postData, commentsTree)
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
