package server

import (
	"net/http"
	"web/internal/helpers/response"
)

func APINotFoundHandler(w http.ResponseWriter, r *http.Request) {
	response.SendErr(w, http.StatusNotFound, "not found")
}
