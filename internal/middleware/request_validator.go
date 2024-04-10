package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	oapimiddleware "github.com/oapi-codegen/nethttp-middleware"

	"github.com/manish-sa/india-template/internal/api/http/oapi"
	"github.com/manish-sa/india-template/internal/logger"
)

func Validator(swagger *openapi3.T) func(next http.Handler) http.Handler {
	return oapimiddleware.OapiRequestValidatorWithOptions(swagger, &oapimiddleware.Options{
		SilenceServersWarning: true,
		ErrorHandler: func(w http.ResponseWriter, message string, statusCode int) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(statusCode)
			response := oapi.ErrorResponse{
				Success: false,
				Message: []string{message},
				Data:    make([]interface{}, 0),
			}
			logger.LogWarn(context.TODO(), "Request validator error", message)
			_ = json.NewEncoder(w).Encode(response)
		},
	})
}
