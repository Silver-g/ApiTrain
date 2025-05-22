package postgres

import "ApiTrain/internal/domain"

func (r *Postgres) IsPostTitleExists(title string) (bool, error) { //возможно тут нужн передать не просто строку а тут сделать экземпляр структуры чтобы в сервисе не передавать поле а передать всю структуру
	var exists bool
	query := "SELECT EXISTS (SELECT 1 FROM posts WHERE title = $1)" //но звучит как хуйня поэтому сделаю как считаю нужным но мб не правильно уточнить У ЮЛИ
	err := r.db.QueryRow(query, title).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil

}
func (r *Postgres) CreatePost(createPostData *domain.CreatePostInternal) (*domain.CreatePostInternal, error) {
	query := "INSERT INTO posts (title, text, comments_enabled, userid) VALUES ($1, $2, $3, $4) RETURNING postid"
	err := r.db.QueryRow(query, createPostData.Title, createPostData.Text, createPostData.CommentsEnabled, createPostData.UserId).Scan(&createPostData.PostId) // Уточнить обязателньый ли аргумент Scan я думаю нет и зачем возвращать id я не хотел жинзь застваила
	if err != nil {                                                                                                                                            //добавить булевое поле в запрос (в баундари и тут)
		return nil, err
	}
	return createPostData, nil
}
