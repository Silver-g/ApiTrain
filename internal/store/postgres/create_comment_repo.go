package postgres

import (
	"ApiTrain/internal/domain"
	"database/sql"
	"errors"
	"fmt"
)

var ErrPostNotFound error = errors.New("post not found")
var ErrCommentsDisabled error = errors.New("сomments are closed for this post")

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
