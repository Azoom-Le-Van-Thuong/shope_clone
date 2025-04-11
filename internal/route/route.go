package route

import (
	"github.com/gin-gonic/gin"
	"shope_clone/internal/container"
	"shope_clone/internal/user/handler"
)

func SetupRoutes(r *gin.Engine, ctn *container.Container) {
	// Register other routes here
	handler.RegisterUserRoutes(r, ctn.CreateUserUC, ctn.GetUserUC)
}
