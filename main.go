package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"web/internal/config"
	"web/internal/modules/auth"
	"web/internal/server"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func run() error {
	// load application configs and variables
	config, err := config.NewConfig()
	if err != nil {
		return err
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// init and test database connection
	db, err := sqlx.Open("postgres", config.Database.ConnectionString)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %s", err.Error())
	}
	defer func() {
		if err := db.Close(); err != nil {
			logger.Error("failed to close database connection", "error", err.Error())
		}
	}()

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to communicate with the database: %s", err.Error())
	}

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
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		os.Exit(1)
	}
}
