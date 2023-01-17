package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

func (ctx Context) Body(target interface{}) error {
	if err := json.NewDecoder(ctx.Request.Body).Decode(target); err != nil {
		return err
	}

	if err := validator.New().Struct(target); err != nil {
		return err
	}

	return nil
}

func (ctx Context) JSON(statusCode int, data interface{}) {
	ctx.Writer.Header().Add("Content-type", "application/json")
	ctx.Writer.WriteHeader(statusCode)
	json.NewEncoder(ctx.Writer).Encode(data)
}

func (ctx Context) Status(statusCode int) {
	ctx.Writer.WriteHeader(statusCode)
}
