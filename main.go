package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"web/internal/config"
	"web/internal/helpers/jwt"
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
	mux.HandleFunc("GET /test", LoggedInMiddleware(config, ProtectedRouteHandler))

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
	}

	response.SendOk(w, token)
}

func ParseBearerToken(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	if header == "" {
		return header, errors.New("missing authorization bearer token")
	}

	if !strings.HasPrefix(header, "Bearer ") {
		return "", errors.New("only bearer tokens are supported for authorization")
	}

	token := strings.Replace(header, "Bearer ", "", 1)
	if len(token) == 0 {
		return "", errors.New("please provide a bearer token for authorization")
	}

	return token, nil
}

func LoggedInMiddleware(config *config.Config, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := ParseBearerToken(r)
		if err != nil {
			response.SendErr(w, http.StatusUnauthorized, err.Error())
		}

		user, err := jwt.ValidateJWT(config.Auth.JWTSecret, token)
		if err != nil {
			response.SendErr(w, http.StatusUnauthorized, err.Error())
		}

		ctx := context.WithValue(r.Context(), "user_id", user.Id)
		ctx = context.WithValue(ctx, "user_role", user.Role)
		next(w, r.WithContext(ctx))
	}
}

func ProtectedRouteHandler(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("user_id").(string)
	if !ok {
		response.SendErr(w, http.StatusUnauthorized, "unauthorized")
	}

	userRole, ok := r.Context().Value("user_role").(string)
	if !ok {
		response.SendErr(w, http.StatusUnauthorized, "unauthorized")
	}

	user := jwt.JWTUser{
		Id:   userId,
		Role: userRole,
	}

	response.SendOk(w, user)
}
