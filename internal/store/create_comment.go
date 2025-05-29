package store

import "ApiTrain/internal/domain"

type CommentRepository interface {
	CreateComment(commentData *domain.CreateCommentInternal) (*domain.CreateCommentInternal, error)
	CommentsAllowed(postId int) error
	GetCommentsByPostID(postId int) ([]*domain.CreateCommentInternal, error)
	GetPostById(postId int) (*domain.PostResponse, error) //мб перенести
}
