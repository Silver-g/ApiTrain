package postrepo

import "ApiTrain/internal/domain"

type CreatePostRepo interface {
	CreatePost(createPostData *domain.CreatePostInternal) (*domain.CreatePostInternal, error)
	IsPostTitleExists(title string) (bool, error)
}
