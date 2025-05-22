package domain

type CreatePostInternal struct {
	Title  string `json:"title"`
	Text   string `json:"text"`
	UserId int    `json:"userid"` // мб не нужна подпись json узер айди потом проверь в конце мб вообще в внутреней это не нужно
	PostId int    `json:"postid"`
}
