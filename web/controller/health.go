package controller

import (
	"net/http"
)

func Health(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte("Healthy"))
	writer.WriteHeader(http.StatusOK)
}
