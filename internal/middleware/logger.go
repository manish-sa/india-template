package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/manish-sa/india-template/internal/config"
	constants "github.com/manish-sa/india-template/internal/constant"
	"github.com/manish-sa/india-template/internal/helper"
	"github.com/manish-sa/india-template/internal/logger"
	"github.com/manish-sa/india-template/internal/model"

	"gitlab.dyninno.net/go-libraries/fluentdlogger/v2"
	"gitlab.dyninno.net/go-libraries/log"
)

func Logger(cfg config.Config) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			lrw := model.LoggingResponseWriter{
				ResponseWriter: w,
			}

			logId := r.Header.Get("X-Log-Id")
			if logId == "" {
				logId = fluentdlogger.GenerateNewLogID(cfg.Logger.Fluentd.ProjectName)
			}
			body, _ := io.ReadAll(r.Body)
			queryParms := r.URL.Query()
			var requestBody map[string]interface{}
			_ = json.Unmarshal(body, &requestBody)
			r.Body = io.NopCloser(bytes.NewBuffer(body))
			sourceId := requestBody["sourceId"]
			sourceName := requestBody["sourceName"]
			if requestBody["sourceId"] == nil {
				sourceId = queryParms["sourceId"]
				sourceName = queryParms["sourceName"]
			}
			requestLogger := log.DefaultLogger().With(
				"log_id", logId,
				"sourceId", sourceId,
				"sourceName", sourceName,
			)
			r = r.WithContext(log.WithContext(r.Context(), requestLogger))
			w.Header().Add("X-Log-Id", logId)

			urlString := r.URL.String()

			logData := logger.LogData{
				RequestBody: requestBody,
				Params:      queryParms,
				PathParams:  helper.ExtractPathParams(r.URL.Path),
				Method:      r.Method,
				URL:         urlString,
			}
			ctx := context.WithValue(r.Context(), constants.LogDataContextKey, logData)

			next.ServeHTTP(&lrw, r.WithContext(ctx))
		})
	}
}
