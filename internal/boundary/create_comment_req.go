package boundary

import (
	"ApiTrain/internal/domain"
	"errors"
)

type CreateCommentRequest struct {
	ParentId *int   `json:"parentId"` // и в домаин тоже опять в рот его указатель что то связано с преоброзванием нуля уточнить у ЮЛИ типо json не смог прочитать пока не сделал указатель вот ошибка {"error": "InternalError","message": "pq: INSERT или UPDATE в таблице \"comments\" нарушает ограничение внешнего ключа \"comments_parent_id_fkey\"" }
	Text     string `json:"text"`
}

const MaxCommentLenght = 2000

var ErrTooLongСomment error = errors.New("Text is too long: maximum is 2000 characters")
var ErrCannotEmpty error = errors.New("Text cannot be empty")

// validate
func CreateCommentValidate(commentReq CreateCommentRequest) error {
	if commentReq.Text == "" { // чат гпт сказал что не нужно добавлять проверку по полю парант айди я растроился мне кажется логика на сервере должна быть строгой и учитывая вообще все ну ладно пусть умные решают УТОЧНИТЬ У ЮЛИ
		return ErrCannotEmpty
	} // тут сделать валидацию на запрощенные символы

	if len(commentReq.Text) > MaxCommentLenght { //мб есть прекл добавить проверку на уровень бд но мы не пижоны так делать не будем
		return ErrTooLongСomment
	}
	return nil
}

// maping
func CreateCommentMaping(commentReq CreateCommentRequest, userId int, postId int) domain.CreateCommentInternal {
	return domain.CreateCommentInternal{
		PostId:   postId,
		UserId:   userId,
		ParentId: commentReq.ParentId,
		Text:     commentReq.Text,
	}
}
