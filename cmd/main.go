package main

import (
	"shope_clone/internal/bootstrap"
	"shope_clone/internal/container"
	"shope_clone/internal/route"
)

func main() {
	cfg := bootstrap.LoadConfig()
	db := bootstrap.ConnectDB(cfg)

	ctn := container.NewContainer(cfg, db)
	server := bootstrap.NewServer(cfg)

	route.SetupRoutes(server.Engine, ctn)

	server.Run()
}
