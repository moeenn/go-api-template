package auth

import (
	"net/http"
	"web/internal/config"
	"web/internal/helpers/jwt"
	"web/internal/helpers/request"
	"web/internal/helpers/response"
	"web/internal/server/middleware"
)

type AuthController struct {
	Config *config.Config
}

func (c *AuthController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/auth/login", c.Login)
	mux.HandleFunc("GET /api/auth/user", middleware.LoggedInMiddleware(c.Config.Auth.JWTSecret, c.GetUser))
	mux.HandleFunc("GET /api/auth/refresh", middleware.LoggedInMiddleware(c.Config.Auth.JWTSecret, c.IssueRefreshToken))
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

func (c *AuthController) IssueRefreshToken(w http.ResponseWriter, r *http.Request) {
	user, err := request.CurrentUser(r)
	if err != nil {
		response.SendErr(w, http.StatusUnauthorized, "please login to continue")
		return
	}

	// TODO: perform checks to see if user is still valid

	token, err := jwt.NewJWT(c.Config.Auth.JWTSecret, user)
	if err != nil {
		response.SendErr(w, http.StatusUnauthorized, "cannot issue refresh token")
		return
	}

	response.SendOk(w, token)
}
