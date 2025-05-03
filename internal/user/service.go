package user

import (
	"context"
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	DB *sql.DB
}

func (s *Service) Register(ctx context.Context, req RegisterRequest) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = s.DB.ExecContext(ctx, `
		INSERT INTO users (username, email, password_hash)
		VALUES ($1, $2, $3)
	`, req.Username, req.Email, string(hashed))
	return err
}

func (s *Service) AuthenticateUser(ctx context.Context, req LoginRequest) (*User, error) {
	row := s.DB.QueryRowContext(ctx, `SELECT id, username, email, password_hash, role, created_at FROM users WHERE email=$1`, req.Email)

	var u User
	err := row.Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.Role, &u.CreatedAt)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	return &u, nil
}
