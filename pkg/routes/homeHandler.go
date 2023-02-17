package routes

import (
	"net/http"
	"sandbox/pkg/server"
)

type HomeResponse struct {
	Message string `json:"message"`
}

// TODO: handlers should return the response
func HomeHandler(ctx *server.Context) {
	res := HomeResponse{
		Message: "welcome to the home page",
	}

	ctx.JSON(http.StatusOK, res)
}
