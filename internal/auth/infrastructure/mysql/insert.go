package authSql

import (
	"context"
	"shope_clone/internal/auth/domain/entity"
)

func (r *authSQlRepository) Insert(ctx context.Context, user entity.UserCredential) error {

	userCredential := map[string]string{
		"email":    user.Email.Value(),
		"password": user.Password.Hashed(),
	}
	if err := r.db.WithContext(ctx).Create(&userCredential).Error; err != nil {
		return err
	}
	return nil
}
