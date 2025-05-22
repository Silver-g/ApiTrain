package boundary

import (
	"ApiTrain/internal/domain"
	"errors"
)

type CreatePostRequest struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func CreatePostPostValidate(createPostReq CreatePostRequest) error {
	if createPostReq.Title == "" || createPostReq.Text == "" {
		return errors.New("Login and password cannot be empty")
	} // тут сделать валидацию на запрощенные символы в лоигне вывести функциив общий файл и все по людски импортировать
	return nil
}
func CreapePostMaping(CreatePostReq CreatePostRequest, userId int) domain.CreatePostInternal {
	return domain.CreatePostInternal{
		Title:  CreatePostReq.Title,
		Text:   CreatePostReq.Text,
		UserId: userId,
	}
}
