package domain

import (
	"context"
	"shope_clone/internal/user/domain/view"
)

type LoginUserRepository interface {
	FindUserByEmail(ctx context.Context, email string) (*view.SimpleUser, error)
}
