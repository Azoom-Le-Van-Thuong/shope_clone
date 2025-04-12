package container

import (
	authSql "shope_clone/internal/auth/infrastructure/mysql"
	authUseCase "shope_clone/internal/auth/usecase"
	"shope_clone/internal/bootstrap"
	userSql "shope_clone/internal/user/infrastructure/mysql"
	userUseCase "shope_clone/internal/user/usecase"
)

type Container struct {
	CreateUserUC        *userUseCase.CreateUserUseCase
	GetUserUC           *userUseCase.GetUserByIDUseCase
	RegisterUserUseCase *authUseCase.RegisterUserUseCase
}

func NewContainer(cfg *bootstrap.Config, db *bootstrap.Database) *Container {
	//userRepo := mysqlUser.NewUserMySQLRepository(db.GormDB)
	//authRepo := mysql.NewAuthRepositorydb.GormDB)
	userRepo := userSql.NewUserMySQLRepository(db.GormDB)
	authRepo := authSql.NewAuthRepository(db.GormDB)

	return &Container{
		CreateUserUC:        userUseCase.NewCreateUserUseCase(userRepo),
		GetUserUC:           userUseCase.NewGetUserByIDUseCase(userRepo),
		RegisterUserUseCase: authUseCase.NewRegisterUserUseCase(authRepo),
	}
}
