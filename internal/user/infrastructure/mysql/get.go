package userSql

import (
	"context"
	"shope_clone/internal/user/domain"
)

func (r *userMySQLRepository) GetByID(ctx context.Context, id uint) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
