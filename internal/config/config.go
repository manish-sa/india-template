package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

const (
	EnvProduction = "production"
	EnvStaging    = "staging"
	EnvLocal      = "local"
)

type (
	Config struct {
		App       AppConfig
		Ports     PortsConfig
		Logger    LoggerConfig
		Sentry    SentryConfig
		Tracing   TracingConfig
		GBO       GBOConfig
		Dreampass DreampassGWConfig
	}

	GBOConfig struct {
		BaseURL string `envconfig:"GBO_CLIENT_COMPONENT_URL"`
	}

	AppConfig struct {
		Env  string `envconfig:"APP_ENV" default:"local"`
		Name string `envconfig:"APP_NAME"`
	}

	PortsConfig struct {
		HTTP uint `envconfig:"PORT_HTTP" default:"80"`
		GRPC uint `envconfig:"PORT_GRPC" default:"82"`
	}

	LoggerConfig struct {
		Level   string `envconfig:"LOG_LEVEL" default:"info"`
		Output  string `envconfig:"LOG_OUTPUT" default:"fluentd"`
		Fluentd struct {
			Host        string `envconfig:"LOG_FLUENTD_HOST" default:"localhost"`
			Port        int    `envconfig:"LOG_FLUENTD_PORT" default:"24224"`
			ProjectName string `envconfig:"LOG_FLUENTD_PROJECT"`
		}
	}

	DreampassGWConfig struct {
		DreampassHTTPTimeout   uint64 `envconfig:"DREAMPASS_GW_HTTP_TIMEOUT" default:"30"`
		DreampassHTTPKeepAlive uint64 `envconfig:"DREAMPASS_GW_HTTP_KEEP_ALIVE"  default:"30"`
		DreampassGWHTTPBaseURL string `envconfig:"DREAMPASS_GW_HTTP_BASE_URL"`
		ServiceName            string `envconfig:"APP_NAME"`
	}

	SentryConfig struct {
		DSN string `envconfig:"SENTRY_DSN"`
	}

	TracingConfig struct {
		Enabled bool    `envconfig:"TRACING_ENABLED" default:"false"`
		Ratio   float64 `envconfig:"TRACING_RATIO" default:"0.1"`
	}
)

func NewConfig() (Config, error) {
	cfg := Config{}

	err := envconfig.Process("", &cfg)
	if err != nil {
		return cfg, errors.Wrap(err, "cannot process the config")
	}

	return cfg, nil
}

func MustNewConfig() Config {
	cfg, err := NewConfig()
	if err != nil {
		panic(err)
	}

	return cfg
}
