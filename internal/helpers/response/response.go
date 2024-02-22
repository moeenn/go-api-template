package response

import (
	"encoding/json"
	"net/http"
)

type apiOkResponse[T any] struct {
	Success bool `json:"success"`
	Data    T    `json:"data"`
}

type apiErrorResponse struct {
	Success    bool   `json:"success"`
	Error      string `json:"error"`
	StatusCode int    `json:"statusCode"`
}

func NewErrorResponse(statusCode int, error string) apiErrorResponse {
	return apiErrorResponse{
		Success:    false,
		Error:      error,
		StatusCode: statusCode,
	}
}

func NewOkResponse[T any](data T) apiOkResponse[T] {
	return apiOkResponse[T]{
		Success: true,
		Data:    data,
	}
}

func SendOk[T any](w http.ResponseWriter, data T) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(NewOkResponse(data))
}

func SendErr(w http.ResponseWriter, statusCode int, error string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(NewErrorResponse(statusCode, error))
}
