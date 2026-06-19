package repository

import (
	"context"

	"github.com/472893749723489727432hjsdjkgf/ai-hack/internal/domain"
)

type Repository interface {
	CreateNewUserDB(ctx context.Context, user *domain.User) error
	CheckExistsUserDB(ctx context.Context, creds *domain.Credentials) (bool, error)
	DeleteUserDB(ctx context.Context, user *domain.User) error
}
