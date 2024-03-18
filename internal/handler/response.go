package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Massage string `json:"message"`
}

func newErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	logrus.Error("Error", message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errorResponse{message})
}
