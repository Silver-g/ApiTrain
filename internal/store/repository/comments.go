package repository

import "ApiTrain/internal/domain"

type CommentRepository interface {
	CreateComment(commentData *domain.CreateCommentInternal) (*domain.CreateCommentInternal, error)
	GetCommentsByPostID(postId int) ([]*domain.CreateCommentInternal, error)
	GetParentExists(parentId int, postId int) error
}
