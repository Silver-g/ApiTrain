package memory

import (
	"ApiTrain/internal/domain"
	"ApiTrain/internal/service/commentservice"
	"sync"
)

type CommentsMemoryRepo struct {
	comments map[int]*domain.CreateCommentInternal
	nextId   int
	mu       sync.RWMutex
}

func NewMemoryCommentRepo() *CommentsMemoryRepo {
	return &CommentsMemoryRepo{
		comments: make(map[int]*domain.CreateCommentInternal),
		nextId:   1,
	}
}

func (m *CommentsMemoryRepo) GetParentExists(parentId int, postId int) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, exists := range m.comments {
		if exists.Id == parentId && exists.PostId == postId {
			return nil
		}
	}
	return commentservice.ErrParentIdNotFound
}

func (m *CommentsMemoryRepo) CreateComment(commentData *domain.CreateCommentInternal) (*domain.CreateCommentInternal, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	commentData.Id = m.nextId
	m.comments[commentData.Id] = commentData
	m.nextId++
	return commentData, nil
}

func (m *CommentsMemoryRepo) GetCommentsByPostID(postId int) ([]*domain.CreateCommentInternal, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	var comments []*domain.CreateCommentInternal
	for _, exists := range m.comments {
		if exists.PostId == postId {
			comments = append(comments, exists)
		}
	}
	return comments, nil
}
