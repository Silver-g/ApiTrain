package handlers

import (
	"ApiTrain/internal/boundary"
	"ApiTrain/internal/handlers/commenthandler"
	"ApiTrain/internal/handlers/posthandler"
	"net/http"
	"strings"
)

type RouteInfo struct {
	CreatePostHandler            *posthandler.HandlerCreatePost
	GetALLPostsHandler           *posthandler.HandlerGetAllPosts
	UpdateCommentsEnabledHandler *posthandler.HandlerUpdateCommentsEnabled
	//comments
	CreateHandler *commenthandler.HandlerCreateComment
	TreeHandler   *commenthandler.HandlerGetPostComments
}

func (ri *RouteInfo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	spliUrl := strings.TrimPrefix(r.URL.Path, "/posts")
	spliUrl = strings.Trim(spliUrl, "/") // убираем слеши с краёв

	var parts []string
	if spliUrl != "" {
		parts = strings.Split(spliUrl, "/")
	}
	//posts работа с постами
	if r.Method == http.MethodPost && len(parts) == 0 {
		ri.CreatePostHandler.CreatePostHandler(w, r)
		return
	}
	if r.Method == http.MethodGet && len(parts) == 0 {
		ri.GetALLPostsHandler.GetAllPostsHandler(w, r)
		return
	}
	if r.Method == http.MethodPatch && len(parts) == 2 && parts[1] == "comments-enabled" {
		ri.UpdateCommentsEnabledHandler.UpdateCommentsEnabled(w, r)
		return
	}
	//comments работа с комментариями
	if r.Method == http.MethodPost && len(parts) == 2 && parts[1] == "comments" {
		ri.CreateHandler.CreateCommentHandler(w, r)
		return
	}
	if r.Method == http.MethodGet && len(parts) == 2 && parts[1] == "comments" {
		ri.TreeHandler.BuildTreeHandler(w, r)
		return
	}
	boundary.WriteResponseErr(w, 404, boundary.ErrorResponse{
		ErrorCode: "NotFound",
		Message:   "Такого маршрута не существует",
	})
}
