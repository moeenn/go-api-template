package server

import (
	"net/http"
)

type Server struct {
	Routes map[string]RouteHandler
}

func (server Server) Run(port string) {
	http.HandleFunc("/", Router(server.Routes))
	http.ListenAndServe(port, nil)
}

func New(routes []Route) *Server {
	routesMap := make(map[string]RouteHandler)

	for _, route := range routes {
		routesMap[route.Key] = route.Handler
	}

	return &Server{
		Routes: routesMap,
	}
}
