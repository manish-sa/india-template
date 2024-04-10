package app

import (
	"context"
	"net/http"

	"github.com/manish-sa/india-template/internal/app/provider"

	apiHttp "github.com/manish-sa/india-template/internal/api/http"

	sentry "github.com/getsentry/sentry-go"
	"github.com/manish-sa/india-template/internal/logger"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"gitlab.dyninno.net/go-libraries/shutdown"
	"gitlab.dyninno.net/go-libraries/tracing"

	"github.com/manish-sa/india-template/internal/config"
	sentryPkg "github.com/manish-sa/india-template/pkg/sentry"
)

type App struct {
	httpSrv *http.Server
	grpcSrv *grpc.Server
}

func MustNew(ctx context.Context, cfg config.Config) *App {
	app, err := New(ctx, cfg)
	if err != nil {
		panic(err)
	}

	return app
}

func New(ctx context.Context, cfg config.Config) (*App, error) {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              cfg.Sentry.DSN,
		EnableTracing:    cfg.Tracing.Enabled,
		Environment:      cfg.App.Env,
		TracesSampleRate: cfg.Tracing.Ratio,
	}); err != nil {
		return nil, errors.Wrap(err, "sentry init")
	}

	shutdown.Add("sentry", func(_ context.Context) error {
		sentryPkg.Flush()
		return nil
	})

	err := tracing.RegisterTracing(ctx, tracing.TracerProviderConfig{
		Enabled:     cfg.Tracing.Enabled,
		ServiceName: cfg.App.Name,
		Env:         tracing.Env(cfg.App.Env),
		Ratio:       &cfg.Tracing.Ratio,
	})
	if err != nil {
		return nil, errors.Wrap(err, "register tracing")
	}

	serviceProviders := provider.NewServiceProviders(ctx, cfg)
	httpSrv := apiHttp.New(cfg, serviceProviders)

	return instance(httpSrv, nil), nil
}

func instance(httpSrv *http.Server, grpcSrv *grpc.Server) *App {
	return &App{
		httpSrv: httpSrv,
		grpcSrv: grpcSrv,
	}
}

func (a *App) Run(ctx context.Context) {
	go func() {
		group, _ := errgroup.WithContext(ctx)

		if a.httpSrv != nil {
			group.Go(func() error {
				logger.LogInfo(ctx, "starting http server", a.httpSrv.Addr)

				if err := a.httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					return errors.Wrap(err, "http server failed")
				}

				return nil
			})
		}

		if a.grpcSrv != nil {
			group.Go(func() error {
				logger.LogInfo(ctx, "starting grpc server", logger.LogData{})
				logger.LogInfo(ctx, "grpc is not implemented", logger.LogData{})

				return nil
			})
		}

		if err := group.Wait(); err != nil {
			// TODO LOG FATAL
			panic(err)
		}
	}()
}
