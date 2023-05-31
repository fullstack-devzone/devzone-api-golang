package posts

import (
	"context"

	"github.com/jackc/pgx/v5"
	log "github.com/sirupsen/logrus"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/users"
)

type PostRepository struct {
	conn *pgx.Conn
}

func NewPostRepo(conn *pgx.Conn) PostRepository {
	return PostRepository{
		conn: conn,
	}
}

func (p PostRepository) GetPosts(ctx context.Context) ([]Post, error) {
	rows, err := p.conn.Query(ctx, `SELECT id, title, url, content, created_at FROM posts`)
	if err != nil {
		return nil, err
	}
	var posts []Post

	defer rows.Close()
	for rows.Next() {
		var post = Post{}
		err = rows.Scan(&post.Id, &post.Title, &post.Url, &post.Content, &post.CreatedDate)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (p PostRepository) GetPostById(ctx context.Context, postId int) (Post, error) {
	log.Infof("Fetching post with id=%d", postId)
	var post = Post{CreatedBy: users.User{}}
	err := p.conn.QueryRow(ctx, `select id, title, url, content, created_by, created_at, updated_at FROM posts where id=$1`, postId).Scan(
		&post.Id, &post.Title, &post.Url, &post.Content, &post.CreatedBy.Id, &post.CreatedDate, &post.UpdatedDate)
	if err != nil {
		return Post{}, err
	}
	return post, nil
}

func (p PostRepository) CreatePost(ctx context.Context, post Post) (Post, error) {
	var lastInsertID int
	err := p.conn.QueryRow(ctx, "insert into posts(title, url, content, created_by, created_at) values($1, $2, $3,$4, $5) RETURNING id",
		post.Title, post.Url, post.Content, post.CreatedBy.Id, post.CreatedDate).Scan(&lastInsertID)
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
