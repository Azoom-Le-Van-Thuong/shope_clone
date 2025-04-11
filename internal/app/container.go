package app

import (
	"gorm.io/gorm"
	"shope_clone/internal/user/application"
	"shope_clone/internal/user/delivery"
	"shope_clone/internal/user/infrastructure"
)

type Container struct {
	UserHandler *delivery.UserHandler
}

func NewContainer(db *gorm.DB) *Container {
	userRepo := infrastructure.NewUserRepository(db)
	userUC := application.NewUserUseCase(userRepo)
	userHandler := delivery.NewUserHandler(userUC)

	return &Container{
		UserHandler: userHandler,
	}
}
