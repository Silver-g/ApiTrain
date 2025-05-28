package store

import "ApiTrain/internal/domain"

type CreateCommentRepo interface {
	CreateComment(commentData *domain.CreateCommentInternal) (*domain.CreateCommentInternal, error)
	CommentsAllowed(postId int) error
}
