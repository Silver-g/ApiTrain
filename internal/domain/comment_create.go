package domain

type CreateCommentInternal struct {
	Id       int    `json:"id"`
	PostId   int    `json:"post_id"`
	UserId   int    `json:"user_id"`
	ParentId *int   `json:"parent_id"`
	Text     string `json:"text"`
}
