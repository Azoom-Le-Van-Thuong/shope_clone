package mysql

import (
	"context"
	"shope_clone/internal/user/domain"
)

func (r *UserMySQLRepository) Insert(ctx context.Context, user *domain.User) error {
	return r.db.Create(user).Error
}
