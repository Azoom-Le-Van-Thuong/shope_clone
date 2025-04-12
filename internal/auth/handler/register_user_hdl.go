package authHandler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shope_clone/internal/auth/dto"
	authUseCase "shope_clone/internal/auth/usecase"
	"shope_clone/pkg/errs"
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
		errs.HandleError(c, err)
		return
	}

	if err := h.UC.Execute(c.Request.Context(), userBody); err != nil {
		errs.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}
