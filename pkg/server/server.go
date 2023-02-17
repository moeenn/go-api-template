package server

import (
	"net/http"
	"sandbox/pkg/log"
)

type Server struct {
	Routes map[string]RouteHandler
	Logger *log.Logger
}

func (server Server) Run(port string) {
	http.HandleFunc("/", Router(server.Routes, server.Logger))

	server.Logger.Log(log.INFO, "starting server on port "+port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		server.Logger.Log(log.ERROR, err.Error())
	}
}

func New(routes []Route, logger *log.Logger) *Server {
	routesMap := make(map[string]RouteHandler)

	for _, route := range routes {
		routesMap[route.Key] = route.Handler
	}

	return &Server{
		Routes: routesMap,
		Logger: logger,
	}
}
