package boundary

import (
	"ApiTrain/internal/domain"
	"errors"
)

type CreatePostRequest struct {
	Title           string `json:"title"`
	Text            string `json:"text"`
	CommentsEnabled *bool  `json:"comments_enabled"` // тут добавил указатель для валидации узнать насколько это адкватно
} // уточнить влияет ли это как то на внутренние уровни или это чисто верхнеуровневая фича
type UpdatePostRequest struct {
	CommentsEnabled *bool `json:"comments_enabled"`
}

func UpdateCommentsEnabledValidate(updateCommentsEnabledReq UpdatePostRequest) error {
	if updateCommentsEnabledReq.CommentsEnabled == nil {
		return errors.New("Selected values cannot be empty")
	}
	return nil
}
func UpdateCommentsEnabledMaping(updateCommentsEnabledReq *UpdatePostRequest, postId int, userId int) domain.UpdatePostRequestInternal {
	return domain.UpdatePostRequestInternal{
		PostId:          postId,
		UserId:          userId,
		CommentsEnabled: *updateCommentsEnabledReq.CommentsEnabled,
	}
}
func CreatePostPostValidate(createPostReq CreatePostRequest) error {
	if createPostReq.Title == "" || createPostReq.Text == "" || createPostReq.CommentsEnabled == nil { // тут добавил проверку мб она избыточна так как фронт все сделает но вообще по хорошему на серваке проверять все
		return errors.New("Text and Title cannot be empty")
	} // тут сделать валидацию на запрощенные символы в лоигне вывести функции в общий файл и все по людски импортировать
	return nil
}
func CreapePostMaping(createPostReq *CreatePostRequest, userId int) domain.CreatePostInternal {
	return domain.CreatePostInternal{
		Title:           createPostReq.Title,
		Text:            createPostReq.Text,
		UserId:          userId,
		CommentsEnabled: *createPostReq.CommentsEnabled,
	}
}
