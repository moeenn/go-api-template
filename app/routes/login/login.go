package login

import (
	"fmt"
	"net/http"
	"sandbox/pkg/server"
)

type Response struct {
	Name  string `json:"name" validate:"required,min=3"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"required,min=18,max=100"`
}

func LoginHandler(ctx *server.Context) error {
	res := &Response{}

	if err := ctx.Body(res); err != nil {
		ctx.Status(http.StatusUnprocessableEntity)

		// TODO: send validation error to client
		fmt.Printf("%v\n", err.Error())
		return nil
	}

	return ctx.JSON(http.StatusOK, res)
}
