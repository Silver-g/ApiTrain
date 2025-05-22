package postService

import (
	"ApiTrain/internal/domain"
	"ApiTrain/internal/store"
)

type CreatePost interface {
	PostCreate(createPostData domain.CreatePostInternal) (int, error)
}

type CreatePostService struct {
	createPostRepo store.CreatePostRepo
}

func CreatePostServiceConstruct(postCreateRepo store.CreatePostRepo) *CreatePostService {
	var createPostServicePointer CreatePostService
	createPostServicePointer.createPostRepo = postCreateRepo
	return &createPostServicePointer
}

func (s *CreatePostService) PostCreate(createPostMapData domain.CreatePostInternal) (int, error) {
	postId, err := s.createPostRepo.CreatePost(&createPostMapData)
	if err != nil {
		return 0, err
	}
	return postId.PostId, nil
}
