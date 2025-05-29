package postgres

import (
	"ApiTrain/internal/domain"
	"database/sql"
	"errors"
	"fmt"
)

var ErrPostNotFound error = errors.New("post not found")
var ErrCommentsDisabled error = errors.New("сomments are closed for this post")

func (r *Postgres) GetPostById(postId int) (*domain.PostResponse, error) {
	var postResponseData domain.PostResponse
	query := "SELECT id, title, text FROM posts WHERE id = $1"
	err := r.db.QueryRow(query, postId).Scan(&postResponseData.Id, &postResponseData.Title, &postResponseData.Text)
	if err != nil {
		return nil, err
	}
	return &postResponseData, nil
} //мб перести

func (r *Postgres) GetCommentsByPostID(postId int) ([]*domain.CreateCommentInternal, error) {
	var err error
	var comments []*domain.CreateCommentInternal //повторить синтаксис срезов
	rows, err := r.db.Query("SELECT id, post_id, user_id, parent_id, text FROM comments WHERE post_id = $1", postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var comment domain.CreateCommentInternal
		var parentID sql.NullInt64                                                               // конеспект про такой тип данных
		err = rows.Scan(&comment.Id, &comment.PostId, &comment.UserId, &parentID, &comment.Text) //напомнить что возвращает скан забыл обработать ошибку
		if err != nil {
			return nil, err
		}
		if parentID.Valid {
			id := int(parentID.Int64) //законспектировать повторить теорию
			comment.ParentId = &id
		} else {
			comment.ParentId = nil
		}
		comments = append(comments, &comment)
	}
	err = rows.Err() //зачем вообще такая ошибка конспект + конспкт по Get
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *Postgres) CommentsAllowed(postId int) error {
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

func (r *Postgres) CreateComment(commentData *domain.CreateCommentInternal) (*domain.CreateCommentInternal, error) {
	var err error                                                                                            //вынести глобально?
	query := "INSERT INTO comments (post_id, user_id, parent_Id, text) VALUES ($1, $2, $3, $4) RETURNING id" // у постов пофикси id
	err = r.db.QueryRow(query, commentData.PostId, commentData.UserId, commentData.ParentId, commentData.Text).Scan(&commentData.Id)
	if err != nil {
		return nil, err
	}
	return commentData, err
}
