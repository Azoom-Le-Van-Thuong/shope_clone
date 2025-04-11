package app

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"shope_clone/pkg/database"
)

func Bootstrap() error {
	_ = godotenv.Load()

	db := database.NewMySQL()
	r := gin.Default()

	c := NewContainer(db)
	RegisterRoutes(r, c)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return r.Run(":" + port)
}
