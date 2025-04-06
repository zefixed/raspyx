package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"raspyx/internal/domain/models"
	"raspyx/internal/repository"
	"strings"
)

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	const op = "repository.postgres.UserRepository.Create"

	query := `INSERT INTO users (uuid, username, password_hash, access_level)
			  VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(ctx, query, user.UUID, user.Username, user.PasswordHash, user.AccessLevel)
	if err != nil {
		if strings.Contains(err.Error(), "23505") {
			return fmt.Errorf("%s: %w", op, repository.ErrExist)
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *UserRepository) Get(ctx context.Context) ([]*models.User, error) {
	const op = "repository.postgres.UserRepository.Get"

	query := `SELECT uuid, username, password_hash, access_level
			  FROM users`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var users []*models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.UUID, &user.Username, &user.PasswordHash, &user.AccessLevel)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		users = append(users, &user)
	}

	return users, nil
}

func (r *UserRepository) GetByUUID(ctx context.Context, uuid uuid.UUID) (*models.User, error) {
	const op = "repository.postgres.UserRepository.GetByUUID"

	query := `SELECT uuid, username, password_hash, access_level
			  FROM users
			  WHERE uuid = $1`
	row := r.db.QueryRow(ctx, query, uuid)
	var user models.User
	err := row.Scan(&user.UUID, &user.Username, &user.PasswordHash, &user.AccessLevel)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &user, nil
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	const op = "repository.postgres.UserRepository.GetByUsername"

	query := `SELECT uuid, username, password_hash, access_level
			  FROM users
			  WHERE username = $1`
	row := r.db.QueryRow(ctx, query, username)
	var user models.User
	err := row.Scan(&user.UUID, &user.Username, &user.PasswordHash, &user.AccessLevel)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &user, nil
}

func (r *UserRepository) GetByAccessLevel(ctx context.Context, accessLevel int) ([]*models.User, error) {
	const op = "repository.postgres.UserRepository.AccessLevel"

	query := `SELECT uuid, username, password_hash, access_level
			  FROM users
			  WHERE access_level <= $1`
	rows, err := r.db.Query(ctx, query, accessLevel)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var users []*models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.UUID, &user.Username, &user.PasswordHash, &user.AccessLevel)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		users = append(users, &user)
	}

	if len(users) == 0 {
		return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return users, nil
}

func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
	const op = "repository.postgres.UserRepository.Update"

	query := `UPDATE users
			  SET access_level = $1
			  WHERE uuid = $2`
	result, err := r.db.Exec(ctx, query, user.AccessLevel, user.UUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return nil
}

func (r *UserRepository) Delete(ctx context.Context, uuid uuid.UUID) error {
	const op = "repository.postgres.UserRepository.Delete"

	query := `DELETE FROM users WHERE uuid = $1`
	result, err := r.db.Exec(ctx, query, uuid)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return nil
}
