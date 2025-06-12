package memory

import (
	"ApiTrain/internal/domain"
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

func (m *CommentsMemoryRepo) CreateComment(commentData *domain.CreateCommentInternal) (*domain.CreateCommentInternal, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	commentData.Id = m.nextId
	m.comments[commentData.Id] = commentData

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
func (m *CommentsMemoryRepo) GetPostById(postId int) (*domain.PostResponse, error) {
	// Вариант 1: вернуть ошибку (не найден) или
	// Вариант 2: вернуть базовую информацию о посте из memory (см. MemoryPostRepo)
	// Для простоты пока возвращаем PostResponse c Id (Title/Text пустые)
	return &domain.PostResponse{Id: postId, Title: "", Text: ""}, nil
}
func (m *CommentsMemoryRepo) CommentsAllowed(postId int) error {
	// В памяти всегда разрешаем комментировать (нет поля comments_enabled)
	return nil
}
