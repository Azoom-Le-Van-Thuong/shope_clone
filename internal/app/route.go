package app

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, c *Container) {
	api := r.Group("/api/v1")
	{
		c.UserHandler.RegisterRoutes(api)
	}
}
