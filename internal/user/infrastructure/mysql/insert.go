package userSql

import (
	"context"
	"shope_clone/internal/user/domain"
)

func (r *userMySQLRepository) Insert(ctx context.Context, user *domain.User) error {
	return r.db.Create(user).Error
}
