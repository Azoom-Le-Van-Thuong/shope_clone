package userSql

import (
	"gorm.io/gorm"
)

type userMySQLRepository struct {
	db *gorm.DB
}

func NewUserMySQLRepository(db *gorm.DB) *userMySQLRepository {
	return &userMySQLRepository{db: db}
}
