package commentrepo

import (
	"ApiTrain/internal/domain"
	"ApiTrain/internal/service/commentservice"
	"database/sql"
)

type CommentPostgres struct { //сделал так как у умного дяди но мб вообще нужно иначе ЮЛЯ В ПОМОЩЬ
	db *sql.DB
}

func NewPostgresComment(dataBase *sql.DB) *CommentPostgres {
	var CommentRepoPointer CommentPostgres
	CommentRepoPointer.db = dataBase
	return &CommentRepoPointer
}
func (r *CommentPostgres) GetParentExists(parentId int, postId int) error {
	var exists bool
	query := "SELECT EXISTS (SELECT 1 FROM comments WHERE id = $1 AND post_id = $2)"
	err := r.db.QueryRow(query, parentId, postId).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return commentservice.ErrParentIdNotFound //потом когда ошибки вынесешь в отедльный файлик переделать
	}
	return nil
}
func (r *CommentPostgres) GetCommentsByPostID(postId int) ([]*domain.CreateCommentInternal, error) {
	var err error
	var comments []*domain.CreateCommentInternal //повторить синтаксис срезов
	query := "SELECT id, post_id, user_id, parent_id, text FROM comments WHERE post_id = $1"
	rows, err := r.db.Query(query, postId)
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
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *CommentPostgres) CreateComment(commentData *domain.CreateCommentInternal) (*domain.CreateCommentInternal, error) {
	var err error                                                                                            //вынести глобально?
	query := "INSERT INTO comments (post_id, user_id, parent_Id, text) VALUES ($1, $2, $3, $4) RETURNING id" // у постов пофикси id
	err = r.db.QueryRow(query, commentData.PostId, commentData.UserId, commentData.ParentId, commentData.Text).Scan(&commentData.Id)
	if err != nil {
		return nil, err
	}
	return commentData, err
}
