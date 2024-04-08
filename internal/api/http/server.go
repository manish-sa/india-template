package http

import (
    "context"
    "fmt"
    "net/http"
    "time"
    
    "gitlab.dyninno.net/trevolution/ancillaries/lbp/lbc-service/internal/app/provider"
    
    "github.com/pkg/errors"
    
    "gitlab.dyninno.net/go-libraries/shutdown"
    
    "gitlab.dyninno.net/trevolution/ancillaries/lbp/lbc-service/internal/config"
)

const (
    defaultReadTimeout  = 5 * time.Second
    defaultWriteTimeout = 30 * time.Second
)

func New(cfg config.Config, serviceProviders *provider.ServiceProvider) *http.Server {
    api := NewAPI(cfg, serviceProviders)
    
    return NewHTTPServer(cfg.Ports.HTTP, api.NewRouter())
}

func NewHTTPServer(port uint, handler http.Handler) *http.Server {
    addr := fmt.Sprintf(":%d", port)
    
    srv := &http.Server{
        Addr:         addr,
        Handler:      handler,
        ReadTimeout:  defaultReadTimeout,
        WriteTimeout: defaultWriteTimeout,
    }
    
    shutdown.Add(fmt.Sprintf("http-server:%d", port), func(ctx context.Context) error {
        srv.SetKeepAlivesEnabled(false)
        return errors.Wrap(srv.Shutdown(ctx), "http server shutdown failed")
    })
    
    return srv
}
