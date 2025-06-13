package commentservice

import (
	"ApiTrain/internal/domain"
	"ApiTrain/internal/store/postgres/postrepo"
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
	postRepo          repository.CreatePostRepo
	userRepo          repository.UserRepository
}

func NewCommentService(commentRepo repository.CommentRepository, postRepo repository.CreatePostRepo, userRepo repository.UserRepository) *CreateCommentService {
	return &CreateCommentService{
		createCommentRepo: commentRepo, //тут подробно разобрать
		postRepo:          postRepo,
		userRepo:          userRepo,
	}
}
func (s *CreateCommentService) CommentCreate(commentData *domain.CreateCommentInternal) (int, error) { //опять ебучие указатели врот их наоборот//
	err := s.userRepo.GetUserById(commentData.UserId)
	if err != nil {
		return 0, err
	}
	err = s.postRepo.CommentsAllowed(commentData.PostId)
	if err == postrepo.ErrCommentsDisabled {
		return 0, ErrCommentsDisabled
	}
	if err == postrepo.ErrPostNotFound {
		return 0, ErrPostNotFound
	}
	if err != nil {
		return 0, err //не уверен что это надо но осавил так как мб внутреннаяя ошибка в том методе прокнет а так я бы перезаписал err
	}
	if commentData.ParentId != nil {
		err = s.createCommentRepo.GetParentExists(*commentData.ParentId, commentData.PostId)
		if err != nil {
			return 0, err
		}
	}
	commentId, err := s.createCommentRepo.CreateComment(commentData)
	if err != nil {
		return 0, err //можно ли вообще делать return 0
	}
	return commentId.Id, nil
}
