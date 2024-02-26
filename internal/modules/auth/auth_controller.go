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

// allow users to login to the system and receive auth token
func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	// TODO: read request body and validate
	user := jwt.JWTUser{
		Id:   "A100",
		Role: "ADMIN",
	}

	token, err := jwt.NewJWT(c.Config.Auth.JWTSecret, c.Config.Auth.JWTExpiryHours, user)
	if err != nil {
		response.SendErr(w, http.StatusUnauthorized, err.Error())
		return
	}

	response.SendOk(w, token)
}

// return details of the currently logged-in user
func (c *AuthController) GetUser(w http.ResponseWriter, r *http.Request) {
	user, err := request.CurrentUser(r)
	if err != nil {
		response.SendErr(w, http.StatusUnauthorized, err.Error())
		return
	}

	response.SendOk(w, user)
}

// allow logged-in users to get a refreshed auth token in exchange for a
// previously issued (un-expired) token.
func (c *AuthController) IssueRefreshToken(w http.ResponseWriter, r *http.Request) {
	user, err := request.CurrentUser(r)
	if err != nil {
		response.SendErr(w, http.StatusUnauthorized, "please login to continue")
		return
	}

	// TODO: perform checks to see if user is still valid

	token, err := jwt.NewJWT(c.Config.Auth.JWTSecret, c.Config.Auth.JWTExpiryHours, user)
	if err != nil {
		response.SendErr(w, http.StatusUnauthorized, "cannot issue refresh token")
		return
	}

	response.SendOk(w, token)
}
