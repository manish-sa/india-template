package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	chi "github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestValidatorWithSimpleSpecification(t *testing.T) {
	t.Parallel()

	swagger := &openapi3.T{
		OpenAPI: "3.0.0",
		Paths: openapi3.Paths{
			"/test": &openapi3.PathItem{
				Get: &openapi3.Operation{
					Parameters: openapi3.Parameters{},
				},
			},
		},
	}

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r := chi.NewRouter()
	r.Use(Validator(swagger))
	r.Handle("/test", testHandler)

	// missed some_param for test
	request, _ := http.NewRequestWithContext(context.Background(), "GET", "/test", nil)

	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	// Assert
	assert.Equal(t, http.StatusOK, response.Code, "Expected status code 200")
}
