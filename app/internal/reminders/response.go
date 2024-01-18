package reminders

import (
	"encoding/json"
	"net/http"
)

type GenericResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func NewGenericResponse(success bool, message string, data any) GenericResponse {
	return GenericResponse{
		Success: success,
		Message: message,
		Data:    data,
	}
}

func RespondToWriter(writer http.ResponseWriter, success bool, message string, data any, status int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	err := json.NewEncoder(writer).Encode(NewGenericResponse(success, message, data))
	if err != nil {
		_ = json.NewEncoder(writer).Encode(NewGenericResponse(false, "Internal Server Error While Responding", nil))
	}

}
