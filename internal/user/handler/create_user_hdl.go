package handler

import (
	"net/http"
	"shope_clone/internal/user/domain"
	"shope_clone/internal/user/usecase"

	"github.com/gin-gonic/gin"
)

type CreateUserHandler struct {
	UC *usecase.CreateUserUseCase
}

func NewCreateUserHandler(uc *usecase.CreateUserUseCase) *CreateUserHandler {
	return &CreateUserHandler{UC: uc}
}

func (h *CreateUserHandler) Handle(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	if err := h.UC.Execute(c.Request.Context(), &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, user)
}
