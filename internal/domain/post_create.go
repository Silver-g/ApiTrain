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
	Comments []*CommentTree `json:"comments"`
}
