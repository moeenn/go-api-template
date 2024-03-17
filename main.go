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
		exit(err)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// instantiate all controllers here
	authController := auth.AuthController{Config: config}

	// init mux and register all routes here
	mux := http.NewServeMux()
	authController.RegisterRoutes(mux)

	// register catch-all handlers
	mux.HandleFunc("/api/", server.APINotFoundHandler)

	// start the web server process
	logger.Info("starting web server", "address", config.Server.Address())
	if err := http.ListenAndServe(config.Server.Address(), mux); err != nil {
		exit(err)
	}
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
	os.Exit(1)
}
