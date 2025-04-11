package domain

import "context"

type UserRepository interface {
	Insert(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id uint) (*User, error)
}
