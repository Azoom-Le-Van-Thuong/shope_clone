package authSql

import (
	"context"
	"shope_clone/internal/auth/domain/entity"
)

func (r *authSQlRepository) Insert(ctx context.Context, user *entity.UserAuthCreation) error {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}
