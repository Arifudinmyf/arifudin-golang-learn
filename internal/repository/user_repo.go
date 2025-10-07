package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"arifudin-golang-learn/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(ctx context.Context, u *models.User) (int, error) {
	var id int
	query := `INSERT INTO users (name, email, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id`
	now := time.Now()
	err := r.DB.QueryRowContext(ctx, query, u.Name, u.Email, now, now).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (*models.User, error) {
	q := `SELECT id, name, email, created_at, updated_at FROM users WHERE id=$1`
	u := &models.User{}
	row := r.DB.QueryRowContext(ctx, q, id)
	if err := row.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}
