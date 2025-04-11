package bootstrap

import "github.com/gin-gonic/gin"

type Server struct {
	Engine *gin.Engine
	Port   string
}

func NewServer(cfg *Config) *Server {
	engine := gin.Default()

	return &Server{
		Engine: engine,
		Port:   cfg.Port,
	}
}

func (s *Server) Run() {
	s.Engine.Run(":" + s.Port)
}
