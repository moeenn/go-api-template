package main

import (
	"sandbox/pkg/log"
	"sandbox/pkg/routes"
	"sandbox/pkg/server"
)

const (
	PORT = "5000"
)

func main() {
	logger := log.New()

	routes := []server.Route{
		{Key: "GET /", Handler: routes.HomeHandler},
		{Key: "POST /login", Handler: routes.LoginHandler},
	}

	server := server.New(routes, logger)
	server.Run(PORT)
}
