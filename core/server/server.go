package server

import (
	"app/pkg/config"
	"app/pkg/routes"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
	Config config.ServerConfig
}

func New(config config.ServerConfig) Server {
	router := gin.Default()
	router.SetTrustedProxies(config.TrustedProxies)
	routes.RegisterAPIRoutes(router)

	return Server{
		Router: router,
		Config: config,
	}
}

func (s Server) Start() {
	s.Router.Run(s.Config.Port)
}
