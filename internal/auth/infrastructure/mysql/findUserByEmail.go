package authSql

import (
	"context"
	"shope_clone/internal/user/view"
)

func (r *authSQlRepository) FindUserByEmail(ctx context.Context, email string) (*view.SimpleUser, error) {
	var user view.SimpleUser
	err := r.db.Table(view.SimpleUser{}.TableName()).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
