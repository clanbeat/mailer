package main

import (
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"error_message"`
}

func badRequest(err string) (int, interface{}) {
	return http.StatusBadRequest, errorResponse(err)
}

func notFound(err string) (int, interface{}) {
	return http.StatusNotFound, errorResponse(err)
}

func forbiddenRequest() (int, interface{}) {
	return http.StatusForbidden, errorResponse("you are not allowed to do this")
}

func errorResponse(message string) *ErrorResponse {
	return &ErrorResponse{Message: message}
}
