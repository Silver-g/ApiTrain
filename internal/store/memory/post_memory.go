package memory

import (
	"ApiTrain/internal/domain"
	"errors"
	"sync"
)

type MemoryPostRepo struct { //тут менять
	posts  map[int]*domain.CreatePostInternal
	nextId int
	mu     sync.RWMutex
}

func NewMemoryPostRepo() *MemoryPostRepo {
	return &MemoryPostRepo{
		posts:  make(map[int]*domain.CreatePostInternal),
		nextId: 1,
	}
}
func (m *MemoryPostRepo) CreatePost(createPostData *domain.CreatePostInternal) (*domain.CreatePostInternal, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	createPostData.Id = m.nextId
	m.posts[m.nextId] = createPostData
	m.nextId++

	return createPostData, nil
}
func (m *MemoryPostRepo) IsPostTitleExists(title string) (bool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, exist := range m.posts {
		if exist.Title == title {
			return true, nil
		}
	}
	return false, nil
}
func (m *MemoryPostRepo) GetAllPosts() ([]*domain.PostResponse, error) {
	m.mu.Lock() //опять же мб RLock
	defer m.mu.Unlock()
	var posts []*domain.PostResponse
	if len(m.posts) == 0 { //для пагинации нужно будет конечно обработать более сложные случаи
		return nil, errors.New("no posts found")
	}
	for _, post := range m.posts {
		postResp := &domain.PostResponse{
			Id:    post.Id,
			Title: post.Title,
			Text:  post.Text, //это переписать ручкаи потренироватся не сразу допер как записать в структуру но опять же проблемы на уровне domain
		}
		posts = append(posts, postResp) // грязно из за плохой логики с структурами в иделе переделать но это не точно нужно узнать у кого то
	} // и тутне обрабатывая ошибку если не найден исправить
	return posts, nil
}
func (m *MemoryPostRepo) UpdateCommentsEnabled(reqData *domain.UpdatePostRequestInternal) (*domain.UpdatePostRequestInternal, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, post := range m.posts {
		if reqData.PostId == post.Id && reqData.UserId == post.UserId {
			post.CommentsEnabled = reqData.CommentsEnabled
			result := &domain.UpdatePostRequestInternal{
				PostId:          post.Id,
				UserId:          post.UserId,
				CommentsEnabled: post.CommentsEnabled,
			}
			return result, nil
		}
	}
	return nil, errors.New("post not found or not owned by user") //вынести в переменную
}
