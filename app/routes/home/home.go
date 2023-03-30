package home

import (
	"net/http"
	"sandbox/pkg/server"
)

type Response struct {
	Message string `json:"message"`
}

func HomeHandler(ctx *server.Context) error {
	res := Response{
		Message: "welcome to the home page",
	}

	return ctx.JSON(http.StatusOK, res)
}
