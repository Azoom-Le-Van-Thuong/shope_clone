package domain

import (
	"context"
	"shope_clone/internal/user/simple_user_login"
)

type LoginUserRepository interface {
	FindUserByEmail(ctx context.Context, email string) (*view.SimpleUser, error)
}
