package errs

import (
	"github.com/gin-gonic/gin"
	"net/http"
	authdomain "shope_clone/internal/auth/domain"
)

func handleAuthErrors(c *gin.Context, err error) bool {
	switch err {
	case authdomain.ErrInvalidCredential:
		c.JSON(http.StatusUnauthorized, New(http.StatusUnauthorized, err.Error()))
	case authdomain.ErrTokenExpired:
		c.JSON(http.StatusUnauthorized, New(http.StatusUnauthorized, err.Error()))
	case authdomain.ErrTokenInvalid:
		c.JSON(http.StatusUnauthorized, New(http.StatusUnauthorized, err.Error()))
	default:
		return false
	}
	return true
}
