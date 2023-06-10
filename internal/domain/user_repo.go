package domain

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

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
	err := p.conn.QueryRow(ctx, `select id, name, email, password FROM users where email=$1`, username).Scan(
		&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	if CheckPasswordHash(password, user.Password) {
		return &user, nil
	}
	return nil, errors.New("invalid username and password")
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
