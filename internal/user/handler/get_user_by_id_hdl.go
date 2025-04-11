package handler

import (
	"net/http"
	"shope_clone/internal/user/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetUserByIDHandler struct {
	UC *usecase.GetUserByIDUseCase
}

func NewGetUserByIDHandler(uc *usecase.GetUserByIDUseCase) *GetUserByIDHandler {
	return &GetUserByIDHandler{UC: uc}
}

func (h *GetUserByIDHandler) Handle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	user, err := h.UC.Execute(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
