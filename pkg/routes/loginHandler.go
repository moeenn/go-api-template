package routes

import (
	"net/http"
	"sandbox/pkg/server"
)

type LoginResponse struct {
	Name  string `json:"name" validate:"required,min=3"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"required,min=18,max=100"`
}

func LoginHandler(ctx *server.Context) {
	res := &LoginResponse{}

	if err := ctx.Body(res); err != nil {
		ctx.Status(http.StatusUnprocessableEntity)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
