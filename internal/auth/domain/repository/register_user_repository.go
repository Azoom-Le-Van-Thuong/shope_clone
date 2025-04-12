package domain

import (
	"context"
	"shope_clone/internal/auth/domain/entity"
	"shope_clone/internal/user/view"
)

type RegisterUserRepository interface {
	Insert(ctx context.Context, user entity.UserCredential) error
	FindUserByEmail(ctx context.Context, email string) (*view.SimpleUser, error)
}
