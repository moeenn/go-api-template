package main

import (
	"sandbox/app/routes"
	"sandbox/pkg/log"
	"sandbox/pkg/server"
)

const (
	PORT = "5000"
)

func main() {
	logger := log.New()
	server := server.New(routes.Routes, logger)
	server.Run(PORT)
}
