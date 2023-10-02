package main

import (
	"persserv/internal/app"
	"persserv/internal/config"

	_ "github.com/lib/pq"
)

// @title RSOI Lab01 API
// @version 1.0
// @description API Server for RSOI Lab01

// @host http://212.233.95.128:8080
// @BasePath /
func main() {
	cfg := config.Load()
	app.Run(cfg)
}
