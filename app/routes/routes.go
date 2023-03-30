package routes

import (
	"sandbox/app/routes/home"
	"sandbox/app/routes/login"
	"sandbox/pkg/server"
)

var Routes = []server.Route{
	{Key: "GET /", Handler: home.HomeHandler},
	{Key: "POST /login", Handler: login.LoginHandler},
}
