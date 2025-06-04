package postservice

import (
	"ApiTrain/internal/domain"
	"ApiTrain/internal/store/postgres/postrepo"
	"errors"
)

var ErrAlreadyExist error = errors.New("post with this title already exists")

type CreatePost interface {
	PostCreate(createPostData *domain.CreatePostInternal) (int, error)
}

type CreatePostService struct {
	createPostRepo postrepo.CreatePostRepo
}

func NewPostService(postCreateRepo postrepo.CreatePostRepo) *CreatePostService {
	var createPostServicePointer CreatePostService
	createPostServicePointer.createPostRepo = postCreateRepo
	return &createPostServicePointer
}

func (s *CreatePostService) PostCreate(createPostMapData *domain.CreatePostInternal) (int, error) {
	exists, err := s.createPostRepo.IsPostTitleExists(createPostMapData.Title)
	if err != nil {
		return 0, err
	}
	if exists {
		return 0, ErrAlreadyExist
	}
	postData, err := s.createPostRepo.CreatePost(createPostMapData)
	if err != nil {
		return 0, err
	}
	return postData.Id, nil
}
