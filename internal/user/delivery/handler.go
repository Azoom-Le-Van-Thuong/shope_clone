package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shope_clone/internal/user/application"
	"shope_clone/internal/user/domain"
)

type UserHandler struct {
	UC *application.UserUseCase
}

func NewUserHandler(uc *application.UserUseCase) *UserHandler {
	return &UserHandler{UC: uc}
}

func (h *UserHandler) RegisterRoutes(rg *gin.RouterGroup) {
	user := rg.Group("/users")
	{
		user.GET("/", h.GetUsers)
		user.POST("/", h.CreateUser)
	}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.UC.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.UC.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, user)
}
