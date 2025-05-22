package postgres

import "ApiTrain/internal/domain"

func (r *Postgres) CreatePost(createPostData *domain.CreatePostInternal) (*domain.CreatePostInternal, error) {
	query := "INSERT INTO posts (title, text, userid) VALUES ($1, $2, $3) RETURNING postid"
	err := r.db.QueryRow(query, createPostData.Title, createPostData.Text, createPostData.UserId).Scan(&createPostData.PostId) // Уточнить обязателньый ли аргумент Scan я думаю нет и зачем возвращать id я не хотел жинзь застваила
	if err != nil {
		return nil, err
	}
	return createPostData, nil
}
