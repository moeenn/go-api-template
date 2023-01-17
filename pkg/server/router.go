package server

import (
	"fmt"
	"net/http"
)

type RouteHandler func(ctx *Context)

type Route struct {
	Key     string
	Handler RouteHandler
}

func Router(routes map[string]RouteHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestKey := fmt.Sprintf("%s %s", r.Method, r.URL)
		ctx := &Context{w, r}

		handler := routes[requestKey]
		if handler == nil {
			ctx.Status(http.StatusNotFound)
			return
		}

		handler(ctx)
	}
}
