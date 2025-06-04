package commenthandler

import (
	"ApiTrain/internal/boundary"
	"net/http"
	"strings"
)

type CommentRouter struct {
	CreateHandler *HandlerCreateComment
	TreeHandler   *HandlerGetPostComments
}

func (cr *CommentRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost && strings.HasSuffix(r.URL.Path, "/comments") {
		cr.CreateHandler.CreateCommentHandler(w, r)
		return
	}
	if r.Method == http.MethodGet && strings.HasSuffix(r.URL.Path, "/comments") {
		cr.TreeHandler.BuildTreeHandler(w, r)
		return
	}
	boundary.WriteResponseErr(w, 404, boundary.ErrorResponse{
		ErrorCode: "NotFound",
		Message:   "Такого маршрута не существует",
	})
}
