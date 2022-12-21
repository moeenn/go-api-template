package routes

import (
	"app/pkg/controllers/auth"

	"github.com/gin-gonic/gin"
)

/**
 *  register all request handlers here
 *
 */
func RegisterAPIRoutes(router *gin.Engine) *gin.RouterGroup {
	api := router.Group("/api")
	{
		api.POST("/login", auth.LoginHandler)
	}

	return api
}
