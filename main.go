package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"web/internal/config"
	"web/internal/modules/auth"
	"web/internal/server"
)

/*
# TODO
- [ ] Implement Repos (i.e. containers of db table operations)
- [ ] Implement Services (i.e. containers of business logic)
*/

func main() {
	config, err := config.NewConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// instantiate all controllers here
	authController := auth.AuthController{Config: config}

	// init mux and register all routes here
	mux := http.NewServeMux()
	authController.RegisterRoutes(mux)

	// register catch-all handlers
	mux.HandleFunc("/", server.NotFoundHandler)

	// start the web server process
	logger.Info("starting web server", "address", config.Server.Address())
	if err := http.ListenAndServe(config.Server.Address(), mux); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
