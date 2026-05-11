package store

import (
	"database/sql"
	"context"
	"time"
)
type User struct {
	ID int64 `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	password string `json:"-"`
	// profile_pic string `json:"profile_pic_path`
	CreatedAt time.Time `json:"created_at"`
}
type UserStore struct{
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, user *User) error{
	query := `
	INSERT INTO users (username, email, password) VALUES ($1,$2,$3)
	RETURNING id, created_at `
	err := s.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Email,
		user.password,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)
	if err != nil {
		return err
	}
	

	return nil


}