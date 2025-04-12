package authHandler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shope_clone/internal/auth/dto"
	authUseCase "shope_clone/internal/auth/usecase"
)

type registerUserHandler struct {
	UC *authUseCase.RegisterUserUseCase
}

func NewRegisterUserHandler(uc *authUseCase.RegisterUserUseCase) *registerUserHandler {
	return &registerUserHandler{UC: uc}
}

func (h *registerUserHandler) Handle(c *gin.Context) {
	var userBody dto.RegisterRequest
	if err := c.ShouldBindJSON(&userBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	if err := h.UC.Execute(c.Request.Context(), userBody); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}
