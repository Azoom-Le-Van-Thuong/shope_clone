package domain

import (
	"context"
	"shope_clone/internal/auth/domain/entity"
	"shope_clone/internal/user/domain/view"
)

type RegisterUserRepository interface {
	Insert(ctx context.Context, user *entity.UserAuthCreation) error
	FindUserByEmail(ctx context.Context, email string) (*view.SimpleUser, error)
}
