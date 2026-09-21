package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Chethiya-Bandara/altera/backend/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// CReate user
func (r *UserRepository) create(
	ctx context.Context,
	name string,
	email string,
	passwordHash string,
) (*models.User, error) {
	query := `
		INSERT INTO users(name, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id, name, email, password_hash, created_at
	`
	var user models.User

	err := r.db.QueryRow(
		ctx,
		query,
		name,
		email,
		passwordHash,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to create user: %w", err)
	}

	return &user, nil
}

// Get user by email
func (r *UserRepository) GetByEmail(
	ctx context.Context,
	email string,
) (*models.User, error) {
	query := `
		SELECT id, name, password_hash, created_at
		FROM users
		WHERE email = $1
	`

	var user models.User

	err := r.db.QueryRow(
		ctx,
		query,
		email,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("Failed to get user by email: %w", err)
	}

	return &user, nil
}

// Get user by ID
func (r *UserRepository) GetByID(
	ctx context.Context,
	id string,
) (*models.User, error) {
	query := `
		SELECT id, name, email, password_hash, created_at
		FROM users
		WHERE id = $1
	`

	var user models.User

	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}

	return &user, nil
}
