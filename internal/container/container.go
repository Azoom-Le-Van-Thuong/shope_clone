package container

import (
	"shope_clone/internal/bootstrap"
	"shope_clone/internal/user/infrastructure/mysql"
	"shope_clone/internal/user/usecase"
)

type Container struct {
	CreateUserUC *usecase.CreateUserUseCase
	GetUserUC    *usecase.GetUserByIDUseCase
}

func NewContainer(cfg *bootstrap.Config, db *bootstrap.Database) *Container {
	userRepo := mysql.NewUserMySQLRepository(db.GormDB)

	return &Container{
		CreateUserUC: usecase.NewCreateUserUseCase(userRepo),
		GetUserUC:    usecase.NewGetUserByIDUseCase(userRepo),
	}
}
