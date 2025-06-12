package postservice

import (
	"ApiTrain/internal/domain"
	"database/sql"
	"errors"
)

var ErrPostNotFoundOrForbidden = errors.New("post not found or access denied")

func (s *CreatePostService) UpdatePostCommentsEnabled(reqData *domain.UpdatePostRequestInternal) (*domain.UpdatePostRequestInternal, error) {

	updatedEntity, err := s.createPostRepo.UpdateCommentsEnabled(reqData)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrPostNotFoundOrForbidden
		}
		return nil, err
	}
	return updatedEntity, nil
}
