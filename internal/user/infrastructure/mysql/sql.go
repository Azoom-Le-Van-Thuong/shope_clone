package mysql

import (
	"gorm.io/gorm"
)

type UserMySQLRepository struct {
	db *gorm.DB
}

func NewUserMySQLRepository(db *gorm.DB) *UserMySQLRepository {
	return &UserMySQLRepository{db: db}
}
