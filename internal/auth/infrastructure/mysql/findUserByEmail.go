package authSql

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"shope_clone/internal/user/simple_user_login"
)

func (r *authSQlRepository) FindUserByEmail(ctx context.Context, email string) (*view.SimpleUser, error) {
	var user view.SimpleUser
	err := r.db.Table(view.SimpleUser{}.TableName()).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("record not found")
		}
		return nil, errors.New("database error")
	}
	return &user, nil
}
