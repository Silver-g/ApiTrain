package postservice

import (
	"ApiTrain/internal/domain"
	"ApiTrain/internal/store/repository"
	"errors"
)

var ErrAlreadyExist error = errors.New("post with this title already exists")

type PostService interface { //обновил имя для чистоты нужно все интерфейсы вынести наверное
	PostCreate(createPostData *domain.CreatePostInternal) (int, error)
	GetAllPostsServ() ([]*domain.PostResponse, error)
	UpdatePostCommentsEnabled(reqData *domain.UpdatePostRequestInternal) (*domain.UpdatePostRequestInternal, error)
}

type CreatePostService struct {
	createPostRepo repository.CreatePostRepo
	userRepo       repository.UserRepository
}

func NewPostService(postCreateRepo repository.CreatePostRepo, userRepo repository.UserRepository) *CreatePostService {
	return &CreatePostService{
		createPostRepo: postCreateRepo,
		userRepo:       userRepo,
	}
}

func (s *CreatePostService) PostCreate(createPostMapData *domain.CreatePostInternal) (int, error) {
	err := s.userRepo.GetUserById(createPostMapData.UserId)
	if err != nil {
		return 0, err
	}
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
