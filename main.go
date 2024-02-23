package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"web/internal/config"
	"web/internal/helpers/jwt"
	"web/internal/helpers/request"
	"web/internal/helpers/response"
	"web/internal/middleware"
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
	mux.HandleFunc("POST /api/auth/login", c.Login)
	mux.HandleFunc("GET /api/auth/user", middleware.LoggedInMiddleware(c.Config.Auth.JWTSecret, c.GetUser))
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	// TODO: read request body and validate
	user := jwt.JWTUser{
		Id:   "A100",
		Role: "ADMIN",
	}

	token, err := jwt.NewJWT(c.Config.Auth.JWTSecret, user)
	if err != nil {
		response.SendErr(w, http.StatusUnauthorized, err.Error())
		return
	}

	response.SendOk(w, token)
}

func (c *AuthController) GetUser(w http.ResponseWriter, r *http.Request) {
	user, err := request.CurrentUser(r)
	if err != nil {
		response.SendErr(w, http.StatusUnauthorized, err.Error())
		return
	}

	response.SendOk(w, user)
}
