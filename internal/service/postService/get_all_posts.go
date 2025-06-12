package postservice

import "ApiTrain/internal/domain"

func (s *CreatePostService) GetAllPostsServ() ([]*domain.PostResponse, error) {
	postsList, err := s.createPostRepo.GetAllPosts()
	if err != nil {
		return nil, err
	}

	return postsList, nil
}
