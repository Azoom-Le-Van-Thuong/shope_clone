package errs

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleError(c *gin.Context, err error) {
	if handleAuthErrors(c, err) {
		return
	}
	// fallback
	c.JSON(http.StatusInternalServerError, New(http.StatusInternalServerError, "internal server error"))
}
