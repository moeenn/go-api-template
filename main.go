package main

import (
	"sandbox/pkg/routes"
	"sandbox/pkg/server"
)

func main() {
	routes := []server.Route{
		{Key: "GET /", Handler: routes.HomeHandler},
		{Key: "POST /login", Handler: routes.LoginHandler},
	}

	server := server.New(routes)
	server.Run(":5000")
}
