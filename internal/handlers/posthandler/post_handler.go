package posthandler

import (
	"ApiTrain/internal/boundary"
	"net/http"
	"strings"
)

type PostsRouter struct {
	CreatePostHandler            *HandlerCreatePost
	GetALLPostsHandler           *HandlerGetAllPosts
	UpdateCommentsEnabledHandler *HandlerUpdateCommentsEnabled
}

func (pr *PostsRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost && strings.HasSuffix(r.URL.Path, "/posts") {
		pr.CreatePostHandler.CreatePostHandler(w, r)
		return
	}
	if r.Method == http.MethodGet && strings.HasSuffix(r.URL.Path, "/posts") {
		pr.GetALLPostsHandler.GetAllPostsHandler(w, r)
		return
	}
	if r.Method == http.MethodPatch && strings.HasSuffix(r.URL.Path, "/posts") {
		pr.UpdateCommentsEnabledHandler.UpdateCommentsEnabled(w, r)
		return
	}
	boundary.WriteResponseErr(w, 404, boundary.ErrorResponse{
		ErrorCode: "NotFound",
		Message:   "Такого маршрута не существует",
	})
}
