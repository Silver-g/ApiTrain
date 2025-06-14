package postrepo

import (
	"ApiTrain/internal/domain"
	"database/sql"
	"errors"
	"fmt"
)

type PostPostgres struct { //сделал так как у умного дяди но мб вообще нужно иначе ЮЛЯ В ПОМОЩЬ
	db *sql.DB
}

var ErrPostNotFound error = errors.New("post not found")
var ErrCommentsDisabled error = errors.New("сomments are closed for this post")

func NewPostgresPost(dataBase *sql.DB) *PostPostgres {
	var PostRepoPointer PostPostgres
	PostRepoPointer.db = dataBase
	return &PostRepoPointer
}

func (r *PostPostgres) GetPostById(postId int) (*domain.PostResponse, error) {
	var postResponseData domain.PostResponse
	query := "SELECT id, title, text FROM posts WHERE id = $1"
	err := r.db.QueryRow(query, postId).Scan(&postResponseData.Id, &postResponseData.Title, &postResponseData.Text)
	if err != nil {
		return nil, err
	}
	return &postResponseData, nil
} //перенес из комментов

func (r *PostPostgres) CommentsAllowed(postId int) error {
	var commentsEnabled bool
	query := "SELECT comments_enabled FROM posts WHERE id = $1"
	err := r.db.QueryRow(query, postId).Scan(&commentsEnabled)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("Post not found with id=%d", postId) //мб потом убать
			return ErrPostNotFound
		}
		return err
	}
	if !commentsEnabled {
		return ErrCommentsDisabled
	}
	return nil
}

func (r *PostPostgres) UpdateCommentsEnabled(reqData *domain.UpdatePostRequestInternal) (*domain.UpdatePostRequestInternal, error) {
	var result domain.UpdatePostRequestInternal
	query := "UPDATE posts SET comments_enabled = $1 WHERE id = $2 AND user_id = $3 RETURNING id, title, text, comments_enabled"
	err := r.db.QueryRow(query, reqData.CommentsEnabled, reqData.PostId, reqData.UserId).Scan(&result.PostId, &result.Title, &result.Text, &result.CommentsEnabled)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *PostPostgres) GetAllPosts() ([]*domain.PostResponse, error) {
	var posts []*domain.PostResponse
	query := "SELECT id, title, text FROM posts"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var post domain.PostResponse
		err = rows.Scan(&post.Id, &post.Title, &post.Text)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostPostgres) IsPostTitleExists(title string) (bool, error) { //возможно тут нужн передать не просто строку а тут сделать экземпляр структуры чтобы в сервисе не передавать поле а передать всю структуру
	var exists bool
	query := "SELECT EXISTS (SELECT 1 FROM posts WHERE title = $1)" //но звучит как хуйня поэтому сделаю как считаю нужным но мб не правильно уточнить У ЮЛИ
	err := r.db.QueryRow(query, title).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil //похоу бул нахрен не нужен можно сделать проверки на этом уровне а возвращать только ошибку как правильно в душе не секу

}
func (r *PostPostgres) CreatePost(createPostData *domain.CreatePostInternal) (*domain.CreatePostInternal, error) {
	query := "INSERT INTO posts (title, text, comments_enabled, user_id) VALUES ($1, $2, $3, $4) RETURNING id"
	err := r.db.QueryRow(query, createPostData.Title, createPostData.Text, createPostData.CommentsEnabled, createPostData.UserId).Scan(&createPostData.Id) // Уточнить обязателньый ли аргумент Scan я думаю нет и зачем возвращать id я не хотел жинзь застваила
	if err != nil {                                                                                                                                        //добавить булевое поле в запрос (в баундари и тут)
		return nil, err
	}
	return createPostData, nil
}
