package domain

type CreatePostInternal struct {
	Title           string `json:"title"`
	Text            string `json:"text"`
	CommentsEnabled bool   `json:"comments_enabled"`
	UserId          int    `json:"userid"` // мб не нужна подпись json узер айди потом проверь в конце мб вообще в внутреней это не нужно
	Id              int    `json:"id"`
}
type PostResponse struct { //мне очередной раз нужен совет касаемо втрукт для ответа и внутренней работы кажись я чуть хуйню сделал
	Id       int            `json:"id"`
	Title    string         `json:"title"`
	Text     string         `json:"text"`
	Comments []*CommentTree `json:"comments,omitempty"` //если поле не используется то его просто нет
}
type UpdatePostRequestInternal struct {
	PostId          int    `json:"post_id"`
	UserId          int    `json:"user_id"`
	Title           string `json:"title"`
	Text            string `json:"text"`
	CommentsEnabled bool   `json:"comments_enabled"`
}
