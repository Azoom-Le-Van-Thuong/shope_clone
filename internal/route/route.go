package route

import (
	"github.com/gin-gonic/gin"
	authHandler "shope_clone/internal/auth/handler"
	"shope_clone/internal/container"
	userHandler "shope_clone/internal/user/handler"
)

func SetupRoutes(r *gin.Engine, ctn *container.Container) {
	// Register other routes here

	userHandler.RegisterUserRoutes(r, ctn.CreateUserUC, ctn.GetUserUC)
	authHandler.RegisterAuthRoutes(r, ctn.RegisterUserUseCase)

}
