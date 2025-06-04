package commentservice

import (
	"ApiTrain/internal/domain"
	"errors"
)

type CommentService interface {
	GetCommentsData(postId int) ([]*domain.CreateCommentInternal, error)
	GetPostByIdData(postId int) (*domain.PostResponse, error)
}

var ErrParentIdNotFound error = errors.New("parent comment with ID not found")

func CreateResponsePostAndComments(postData *domain.PostResponse, commentsData []*domain.CommentTree) (*domain.PostResponse, error) {
	response := domain.PostResponse{
		Id:       postData.Id,
		Title:    postData.Title,
		Text:     postData.Text,
		Comments: commentsData,
	}
	return &response, nil

}

// надо куда то перенести маперы
// -----------------------------------------------------
func (s *CreateCommentService) GetPostByIdData(postId int) (*domain.PostResponse, error) {
	postData, err := s.createCommentRepo.GetPostById(postId)
	if err != nil {
		return nil, err
	}
	return postData, nil
}

// ////////////////////////////////////////////////
func CreateTreeComments(commentsData []*domain.CreateCommentInternal) ([]*domain.CommentTree, error) { //пока не понятно с указателями можно ли без них кароче сложно сказать почему именно использовал указатель
	valueMap := make(map[int]*domain.CommentTree)
	for _, comment := range commentsData {
		//var node &domain.CommentTree
		node := &domain.CommentTree{ //?? как сдлать более крупную запись
			Id:     comment.Id,
			PostId: comment.PostId,
			UserId: comment.UserId,
			Text:   comment.Text,
		}
		valueMap[comment.Id] = node
	}
	var root []*domain.CommentTree
	for _, comment := range commentsData {
		node := valueMap[comment.Id] // из синтаксиса map получаем тут тело, значение поля мапы с этим айдишником
		if comment.ParentId == nil {
			root = append(root, node)
		} else {
			parent, ok := valueMap[*comment.ParentId]
			if !ok {
				return nil, ErrParentIdNotFound
			}
			parent.Children = append(parent.Children, node)
		}
	}
	return root, nil
}

func (s *CreateCommentService) GetCommentsData(postId int) ([]*domain.CreateCommentInternal, error) {
	var comments []*domain.CreateCommentInternal
	var err error
	comments, err = s.createCommentRepo.GetCommentsByPostID(postId)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
