package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"web/internal/config"
	"web/internal/helpers/response"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// instantiate all controllers here
	authController := AuthController{Config: config}

	// init mux and register all routes here
	mux := http.NewServeMux()
	authController.RegisterRoutes(mux)

	logger.Info("starting web server", "address", config.Server.Address())
	if err := http.ListenAndServe(config.Server.Address(), mux); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}

type AuthController struct {
	Config *config.Config
}

func (c *AuthController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/auth/login", c.Login)
}

type LoginResponse struct {
	Message string `json:"message"`
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	res := LoginResponse{
		Message: "you have logged-in successfully",
	}

	response.SendOk(w, res)
}
