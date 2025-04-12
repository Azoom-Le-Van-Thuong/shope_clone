package authSql

import (
	"context"
	"log"
	"shope_clone/internal/auth/domain/entity"
)

func (r *authSQlRepository) Insert(ctx context.Context, user *entity.UserAuthCreation) error {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		log.Println("Error inserting user:", err)
		return err
	}
	return nil
}
