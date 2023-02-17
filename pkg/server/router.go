package server

import (
	"fmt"
	"net/http"
	"sandbox/pkg/log"
)

type RouteHandler func(ctx *Context)

type Route struct {
	Key     string
	Handler RouteHandler
}

func Router(routes map[string]RouteHandler, logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestKey := fmt.Sprintf("%s %s", r.Method, r.URL)
		ctx := &Context{Writer: w, Request: r}

		handler := routes[requestKey]
		if handler == nil {
			ctx.Status(http.StatusNotFound)
			logger.Log(log.INFO, requestKey+" - 404")
			return
		}

		handler(ctx)

		// TODO: StatusCode not printing properly
		logger.Log(log.INFO, fmt.Sprintf("%s - %d", requestKey, ctx.StatusCode))
	}
}
