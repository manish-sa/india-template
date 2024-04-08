package model

import (
	"fmt"
	"net/http"
)

type (
	LoggingResponseWriter struct {
		http.ResponseWriter
	}
)

func (r *LoggingResponseWriter) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b)
	if err != nil {
		err = fmt.Errorf("error writing response: %w", err)
	}

	return size, err
}

func (r *LoggingResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
}
