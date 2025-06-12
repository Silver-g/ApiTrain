package commentservice

import (
	"ApiTrain/internal/domain"
	"ApiTrain/internal/store/postgres/commentrepo"
	"ApiTrain/internal/store/repository"
	"errors"
)

var ErrPostNotFound error = errors.New("post not found")
var ErrCommentsDisabled error = errors.New("comments are closed for this post")

type CreateCommentServ interface {
	CommentCreate(commentData *domain.CreateCommentInternal) (int, error)
}

type CreateCommentService struct { //переделать название
	createCommentRepo repository.CommentRepository
}

func NewCommentService(commentCreateRepo repository.CommentRepository) *CreateCommentService {
	var CommentCreateServicePointer CreateCommentService
	CommentCreateServicePointer.createCommentRepo = commentCreateRepo
	return &CommentCreateServicePointer
}

func (s *CreateCommentService) CommentCreate(commentData *domain.CreateCommentInternal) (int, error) { //опять ебучие указатели врот их наоборот//
	err := s.createCommentRepo.CommentsAllowed(commentData.PostId)
	if err == commentrepo.ErrCommentsDisabled {
		return 0, ErrCommentsDisabled
	}
	if err == commentrepo.ErrPostNotFound {
		return 0, ErrPostNotFound
	}
	if err != nil {
		return 0, err
	}

	commentId, err := s.createCommentRepo.CreateComment(commentData)
	if err != nil {
		return 0, err //можно ли вообще делать return 0
	}
	return commentId.Id, nil
}
