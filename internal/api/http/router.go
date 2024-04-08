package http

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    
    "gitlab.dyninno.net/go-libraries/fluentdlogger/v2"
    "moul.io/http2curl"
    
    constants "gitlab.dyninno.net/trevolution/ancillaries/lbp/lbc-service/internal/constant"
    "gitlab.dyninno.net/trevolution/ancillaries/lbp/lbc-service/internal/helper"
    internallogger "gitlab.dyninno.net/trevolution/ancillaries/lbp/lbc-service/internal/logger"
    
    chi "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    
    "gitlab.dyninno.net/go-libraries/log"
    "gitlab.dyninno.net/go-libraries/tracing/http/router"
    
    "gitlab.dyninno.net/trevolution/ancillaries/lbp/lbc-service/internal/api/http/oapi"
    "gitlab.dyninno.net/trevolution/ancillaries/lbp/lbc-service/internal/config"
    middle "gitlab.dyninno.net/trevolution/ancillaries/lbp/lbc-service/internal/middleware"
)

type (
    responseMetaData struct {
        status int
        size   int
    }
    
    loggingResponseWriter struct {
        http.ResponseWriter
        responseData *responseMetaData
    }
)

func (api *API) NewRouter() chi.Router {
    r := router.MustNew(router.Config{
        AppName:          api.cfg.App.Name,
        AccessLogEnabled: api.cfg.App.Env != config.EnvProduction,
    })
    
    r.Use(middleware.Recoverer)
    r.Use(logger(api))
    
    if api.cfg.App.Env != config.EnvProduction {
        fs := http.FileServer(http.Dir("./swagger"))
        r.Mount("/swagger/", http.StripPrefix("/swagger/", fs))
    }
    
    fmt.Println("hhhhhhhhhhhhhh")
    
    r.Get("/api/healthcheck", func(w http.ResponseWriter, r *http.Request) {
        response := api.Healthcheck(w, r)
        
        w.Header().Set("Content-Type", "application/json")
        _ = json.NewEncoder(w).Encode(response)
    })
    
    r.Mount("/debug", middleware.Profiler())
    
    r.Group(func(r chi.Router) {
        strictHandler := oapi.NewStrictHandler(api, nil)
        swagger, _ := oapi.GetSwagger()
        r.Use(middle.Validator(swagger))
        oapi.HandlerFromMuxWithBaseURL(strictHandler, r, "/api")
    })
    
    r.Get("/kubernetes/live", api.Live)
    r.Get("/kubernetes/ready", api.Ready)
    
    r.Get("/version", api.Version)
    r.Handle("/metrics", promhttp.Handler())
    
    return r
}

func logger(api *API) func(next http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            responseData := &responseMetaData{
                status: 0,
                size:   0,
            }
            lrw := loggingResponseWriter{
                ResponseWriter: w,
                responseData:   responseData,
            }
            
            logId := r.Header.Get(constants.XLogId)
            if logId == "" {
                logId = fluentdlogger.GenerateNewLogID(api.cfg.Logger.Fluentd.ProjectName)
            }
            
            body, _ := io.ReadAll(r.Body)
            queryParms := r.URL.Query()
            var requestBody map[string]interface{}
            _ = json.Unmarshal(body, &requestBody)
            r.Body = io.NopCloser(bytes.NewBuffer(body))
            
            r.Header.Add(constants.XLogId, logId)
            w.Header().Add(constants.XLogId, logId)
            requestLogger := log.DefaultLogger().With(
                "log_id", logId,
            )
            r = r.WithContext(log.WithContext(r.Context(), requestLogger))
            
            urlString := r.URL.String()
            
            logData := internallogger.LogData{
                RequestBody: requestBody,
                Params:      queryParms,
                PathParams:  helper.ExtractPathParams(r.URL.Path),
                Method:      r.Method,
                URL:         urlString,
            }
            curl, _ := http2curl.GetCurlCommand(r)
            ctx := context.WithValue(r.Context(), constants.Curl, curl)
            ctx = context.WithValue(ctx, constants.LogDataContextKey, logData)
            
            next.ServeHTTP(&lrw, r.WithContext(ctx))
        })
    }
}

func (r *loggingResponseWriter) Write(b []byte) (int, error) {
    size, err := r.ResponseWriter.Write(b)
    r.responseData.size += size
    
    if err != nil {
        err = fmt.Errorf("error writing response: %w", err)
    }
    
    return size, err
}

func (r *loggingResponseWriter) WriteHeader(statusCode int) {
    r.ResponseWriter.WriteHeader(statusCode)
    r.responseData.status = statusCode
}
