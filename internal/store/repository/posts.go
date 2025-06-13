package repository

import "ApiTrain/internal/domain"

type CreatePostRepo interface {
	CreatePost(createPostData *domain.CreatePostInternal) (*domain.CreatePostInternal, error)
	IsPostTitleExists(title string) (bool, error)
	GetAllPosts() ([]*domain.PostResponse, error)
	UpdateCommentsEnabled(reqData *domain.UpdatePostRequestInternal) (*domain.UpdatePostRequestInternal, error)
	//перенес из comments
	GetPostById(postId int) (*domain.PostResponse, error)
	CommentsAllowed(postId int) error
}
