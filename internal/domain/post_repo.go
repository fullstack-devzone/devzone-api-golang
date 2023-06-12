package domain

import (
	"context"
	"fmt"
	"math"
	"strings"

	"github.com/jackc/pgx/v5"
	log "github.com/sirupsen/logrus"
)

type PostRepository struct {
	conn *pgx.Conn
}

func NewPostRepo(conn *pgx.Conn) PostRepository {
	return PostRepository{
		conn: conn,
	}
}

const pageSize = 10

func (p PostRepository) GetPosts(ctx context.Context, page int) (PostsPageModel, error) {
	var totalElements = 0
	err := p.conn.QueryRow(ctx, `SELECT count(*) FROM posts`).Scan(&totalElements)
	if err != nil {
		return PostsPageModel{}, err
	}
	offset := (page - 1) * pageSize
	rows, err := p.conn.Query(ctx,
		`select p.id, p.title, p.url, p.content, p.created_at, p.updated_at,
       				u.id, u.name, u.email
			FROM posts p join users u on p.created_by = u.id
			order by created_at desc OFFSET $1 LIMIT $2`, offset, pageSize)
	if err != nil {
		return PostsPageModel{}, err
	}
	defer rows.Close()
	return buildPagedResult(rows, totalElements, page)
}

func (p PostRepository) SearchPosts(ctx context.Context, keyword string, page int) (PostsPageModel, error) {
	var totalElements = 0
	keyword = fmt.Sprintf("%%%s%%", strings.ToLower(keyword))
	err := p.conn.QueryRow(ctx, `SELECT count(*) FROM posts WHERE lower(title) like $1`, keyword).Scan(&totalElements)
	if err != nil {
		return PostsPageModel{}, err
	}
	offset := (page - 1) * pageSize
	rows, err := p.conn.Query(ctx,
		`select p.id, p.title, p.url, p.content, p.created_at, p.updated_at,
       				u.id, u.name, u.email
			FROM posts p join users u on p.created_by = u.id
			WHERE lower(p.title) like $1
			order by created_at desc OFFSET $2 LIMIT $3`, keyword, offset, pageSize)
	if err != nil {
		return PostsPageModel{}, err
	}
	defer rows.Close()
	return buildPagedResult(rows, totalElements, page)
}

func buildPagedResult(rows pgx.Rows, totalElements, page int) (PostsPageModel, error) {
	var posts []PostModel
	for rows.Next() {
		var post = PostModel{}
		err := rows.Scan(&post.Id, &post.Title, &post.Url, &post.Content, &post.CreatedDate, &post.UpdatedDate,
			&post.CreatedBy.Id, &post.CreatedBy.Name, &post.CreatedBy.Email)
		if err != nil {
			return PostsPageModel{}, err
		}
		posts = append(posts, post)
	}
	totalPages := int(math.Ceil(float64(totalElements) / float64(pageSize)))
	return PostsPageModel{
		TotalElements: totalElements,
		TotalPages:    totalPages,
		PageNumber:    page,
		IsFirst:       page == 1,
		IsLast:        page == totalPages,
		HasNext:       totalPages > page,
		HasPrevious:   page > 1,
		Data:          posts,
	}, nil
}

func (p PostRepository) GetPostById(ctx context.Context, postId int) (PostModel, error) {
	log.Infof("Fetching post with id=%d", postId)
	var post = PostModel{}
	query := `select p.id, p.title, p.url, p.content, p.created_at, p.updated_at,
       				 u.id, u.name, u.email
			  FROM posts p join users u on p.created_by = u.id
			  WHERE p.id=$1`
	err := p.conn.QueryRow(ctx, query, postId).Scan(
		&post.Id, &post.Title, &post.Url, &post.Content, &post.CreatedDate, &post.UpdatedDate,
		&post.CreatedBy.Id, &post.CreatedBy.Name, &post.CreatedBy.Email)
	if err != nil {
		return PostModel{}, err
	}
	return post, nil
}

func (p PostRepository) CreatePost(ctx context.Context, post Post) (Post, error) {
	var lastInsertID int
	err := p.conn.QueryRow(ctx, "insert into posts(title, url, content, created_by, created_at) values($1, $2, $3,$4, $5) RETURNING id",
		post.Title, post.Url, post.Content, post.CreatedBy, post.CreatedDate).Scan(&lastInsertID)
	if err != nil {
		log.Errorf("Error while inserting post row: %v", err)
		return Post{}, err
	}
	post.Id = lastInsertID
	return post, nil
}

func (p PostRepository) UpdatePost(ctx context.Context, post Post) (Post, error) {
	_, err := p.conn.Exec(ctx, "update posts set title = $1, url=$2, content=$3, updated_at=$4 where id=$5",
		post.Title, post.Url, post.Content, post.UpdatedDate, post.Id)
	if err != nil {
		return Post{}, err
	}
	return post, nil
}

func (p PostRepository) DeletePost(ctx context.Context, postId int) error {
	deleteStmt := `delete from posts where id=$1`
	_, err := p.conn.Exec(ctx, deleteStmt, postId)
	return err
}
