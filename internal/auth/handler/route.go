package authHandler

import (
	"github.com/gin-gonic/gin"
	authUseCase "shope_clone/internal/auth/usecase"
)

// impl RegisterAuthRoutes like RegisterUserRoutes

func RegisterAuthRoutes(
	r *gin.Engine,
	registerUC *authUseCase.RegisterUserUseCase,

) {

	registerUserHdl := NewRegisterUserHandler(registerUC)
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", registerUserHdl.Handle)
	}
}
