package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"gitlab.dyninno.net/go-libraries/log"
	"gitlab.dyninno.net/go-libraries/shutdown"

	"gitlab.dyninno.net/trevolution/ancillaries/lbp/lbc-service/internal/app"
	"gitlab.dyninno.net/trevolution/ancillaries/lbp/lbc-service/internal/config"
	"gitlab.dyninno.net/trevolution/ancillaries/lbp/lbc-service/internal/info"
	"gitlab.dyninno.net/trevolution/ancillaries/lbp/lbc-service/pkg/sentry"
)

func apiCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "api server of lbc service",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cfg := config.MustNewConfig()

			logger, err := log.Setup(
				log.WithLogLevel(cfg.Logger.Level),
				log.WithOutput(cfg.Logger.Output),
				log.WithFluentd(
					cfg.Logger.Fluentd.Host,
					cfg.Logger.Fluentd.Port,
					info.GetInfo().ServiceName,
					cfg.Logger.Fluentd.ProjectName,
				),
			)
			if err != nil {
				panic(err)
			}
			defer logger.Close()

			// Use a deferred function to capture panics globally
			defer func() {
				if rec := recover(); rec != nil {
					// Capture the exception with Sentry (if you are using it)
					sentry.Error(ctx, fmt.Sprintf("%v", rec))
					sentry.Flush()
					// Additional cleanup or handling if needed
				}
			}()

			app.MustNew(ctx, cfg).Run(ctx)

			shutdown.Wait(ctx, shutdown.WithLogger(log.DefaultLogger()))
		},
	}
}
