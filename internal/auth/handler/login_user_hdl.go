package authHandler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shope_clone/internal/auth/dto"
	authUseCase "shope_clone/internal/auth/usecase"
	"shope_clone/pkg/errs"
)

type LoginHandler struct {
	UC *authUseCase.LoginUseCase
}

func NewLoginHandler(uc *authUseCase.LoginUseCase) *LoginHandler {
	return &LoginHandler{UC: uc}
}

func (h *LoginHandler) Handle(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errs.HandleError(c, err)
		return
	}

	res, err := h.UC.Execute(c.Request.Context(), req)
	if err != nil {
		errs.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}
