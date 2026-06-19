package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/472893749723489727432hjsdjkgf/ai-hack/internal/domain"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresUserRepository struct {
	db *pgxpool.Pool
}

func NewPostgresUserRepository(db *pgxpool.Pool) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) CreateNewUserDB(ctx context.Context, user *domain.User) error {
	query := `INSERT INTO users (user_name,email,password) VALUES ($1,$2,$3) RETURNING id`
	err := r.db.QueryRow(ctx, query, user.UserName, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return domain.DBErrorUserAlreadyExists
		}
		return fmt.Errorf("postgres: create user failed: %w", err)
	}
	return nil

}

func (r *PostgresUserRepository) CheckExistsUserDB(ctx context.Context, creds *domain.Credentials) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM users WHERE user_name = $1 AND password = $2)`
	var exists bool
	err := r.db.QueryRow(ctx, query, creds.UserName, creds.Password).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("postgres: проверка кредов неудачна: %w", err)
	}
	if !exists {
		return false, domain.DBErrorInvalidCreds
	}
	return true, nil
}

func (r *PostgresUserRepository) DeleteUserDB(ctx context.Context, user *domain.User) error {
	query := `DELETE FROM users WHERE user_name = $1`
	result, err := r.db.Exec(ctx, query, user.UserName)
	if err != nil {
		return fmt.Errorf("postgres: при удалении пользователя ошибка: %w", err)
	}
	if result.RowsAffected() == 0 {
		return domain.DBErrorUserNotFound
	}
	return nil
}
