package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusOk    = "OK"
	StatusError = "Error"
)

func WriteJson(w http.ResponseWriter, statusCode int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) ErrorResponse {
	return ErrorResponse{
		Status: StatusError,
		Error:  err.Error(),
	}
}


func ValidationErrors(errs validator.ValidationErrors) ErrorResponse  {
	var errMsgs []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required" : 
		  	errMsgs = append(errMsgs, fmt.Sprintf("%s field is required", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("%s field is invalid", err.Field()))
		}
	}

	return ErrorResponse{
		Status: StatusError,
		Error: strings.Join(errMsgs, ", "),
	}
}