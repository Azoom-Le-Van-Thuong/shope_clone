package container

import (
	authSql "shope_clone/internal/auth/infrastructure/mysql"
	"shope_clone/internal/auth/token/jwt"
	authUseCase "shope_clone/internal/auth/usecase"
	"shope_clone/internal/bootstrap"
	userSql "shope_clone/internal/user/infrastructure/mysql"
	userUseCase "shope_clone/internal/user/usecase"
)

type Container struct {
	CreateUserUC        *userUseCase.CreateUserUseCase
	GetUserUC           *userUseCase.GetUserByIDUseCase
	RegisterUserUseCase *authUseCase.RegisterUserUseCase
	LoginUserUseCase    *authUseCase.LoginUseCase
}

func NewContainer(cfg *bootstrap.Config, db *bootstrap.Database) *Container {
	//userRepo := mysqlUser.NewUserMySQLRepository(db.GormDB)
	//authRepo := mysql.NewAuthRepositorydb.GormDB)
	userRepo := userSql.NewUserMySQLRepository(db.GormDB)
	authRepo := authSql.NewAuthRepository(db.GormDB)
	jwtProvider := jwt.NewJWTService(cfg.JWTAuthSecret, cfg.JwtAccessTokenExpire)

	return &Container{
		CreateUserUC:        userUseCase.NewCreateUserUseCase(userRepo),
		GetUserUC:           userUseCase.NewGetUserByIDUseCase(userRepo),
		RegisterUserUseCase: authUseCase.NewRegisterUserUseCase(authRepo),
		LoginUserUseCase:    authUseCase.NewLoginUseCase(authRepo, jwtProvider),
	}
}
