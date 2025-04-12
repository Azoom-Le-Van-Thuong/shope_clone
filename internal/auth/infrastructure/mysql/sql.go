package authSql

import "gorm.io/gorm"

type authSQlRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authSQlRepository {
	return &authSQlRepository{db: db}
}
