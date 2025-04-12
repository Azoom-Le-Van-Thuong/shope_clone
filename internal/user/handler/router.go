package userHandler

import (
	"github.com/gin-gonic/gin"
	"shope_clone/internal/user/usecase"
)

func RegisterUserRoutes(
	r *gin.Engine,
	createUC *userUseCase.CreateUserUseCase,
	getUC *userUseCase.GetUserByIDUseCase,

) {
	getUserHandler := NewGetUserByIDHandler(getUC)
	createUserHandler := NewCreateUserHandler(createUC)
	userGroup := r.Group("/users")
	{
		userGroup.POST("", createUserHandler.Handle)
		userGroup.GET("/:id", getUserHandler.Handle)
	}
}
