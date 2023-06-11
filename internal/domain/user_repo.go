package domain

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	log "github.com/sirupsen/logrus"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	conn *pgx.Conn
}

func NewUserRepository(conn *pgx.Conn) UserRepository {
	return UserRepository{
		conn: conn,
	}
}

func (p UserRepository) Login(ctx context.Context, username, password string) (*User, error) {
	var user = User{}
	err := p.conn.QueryRow(ctx, `select id, name, email, password, role FROM users where email=$1`, username).Scan(
		&user.Id, &user.Name, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}
	if CheckPasswordHash(password, user.Password) {
		return &user, nil
	}
	return nil, errors.New("invalid username and password")
}

func (p UserRepository) CreateUser(ctx context.Context, user User) (User, error) {
	var lastInsertID int
	password, err := HashPassword(user.Password)
	if err != nil {
		log.Errorf("Error while encrypting password: %v", err)
		return User{}, err
	}
	err = p.conn.QueryRow(ctx, "insert into users(name, email, password, role, created_at) values($1, $2, $3,$4, $5) RETURNING id",
		user.Name, user.Email, password, user.Role, user.CreatedDate).Scan(&lastInsertID)
	if err != nil {
		log.Errorf("Error while inserting post row: %v", err)
		return User{}, err
	}
	user.Id = lastInsertID
	return user, nil
}

func (p UserRepository) GetUserById(ctx context.Context, userId int) (User, error) {
	log.Infof("Fetching user with id=%d", userId)
	var user = User{}
	err := p.conn.QueryRow(ctx, `select id, name, email, password, role, created_at, updated_at FROM users where id=$1`, userId).Scan(
		&user.Id, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedDate, &user.UpdatedDate)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
