module gitlab.dyninno.net/trevolution/ancillaries/lbp/lbc-service

replace gitlab.dyninno.net/trevolution/ancillaries/lbp/lbc-service => /var/www/html/project

go 1.21.5

require (
	github.com/getkin/kin-openapi v0.120.0
	github.com/getsentry/sentry-go v0.25.0
	github.com/go-chi/chi/v5 v5.0.10
	github.com/joho/godotenv v1.5.1
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/oapi-codegen/nethttp-middleware v1.0.1
	github.com/oapi-codegen/runtime v1.0.0
	github.com/oklog/ulid v1.3.1
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.17.0
	github.com/spf13/cobra v1.8.0
	github.com/stretchr/testify v1.8.4
	gitlab.dyninno.net/go-libraries/client-component-gbo v1.6.0
	gitlab.dyninno.net/go-libraries/client-component-gmail v1.0.2
	gitlab.dyninno.net/go-libraries/dyninnogorm v1.0.0
	gitlab.dyninno.net/go-libraries/fluentdlogger/v2 v2.0.0-alpha.4
	gitlab.dyninno.net/go-libraries/log v1.0.1
	gitlab.dyninno.net/go-libraries/shutdown v1.0.2
	gitlab.dyninno.net/go-libraries/tracing v1.0.5
	golang.org/x/sync v0.3.0
	google.golang.org/grpc v1.58.1
	gorm.io/driver/sqlite v1.5.5
	gorm.io/gorm v1.25.7
)

require (
	github.com/go-ozzo/ozzo-validation/v4 v4.3.0 // indirect
	github.com/smartystreets/goconvey v1.8.1 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.45.0 // indirect
)

require (
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dropbox/godropbox v0.0.0-20220817175148-f0626942059b // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/fluent/fluent-logger-golang v1.9.0 // indirect
	github.com/getsentry/sentry-go/otel v0.25.0 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/swag v0.22.4 // indirect
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/uuid v1.3.1 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/invopop/yaml v0.2.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/perimeterx/marshmallow v1.1.5 // indirect
	github.com/philhofer/fwd v1.1.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.4.1-0.20230718164431-9a2bf3000d16 // indirect
	github.com/prometheus/common v0.44.0 // indirect
	github.com/prometheus/procfs v0.11.1 // indirect
	github.com/riandyrn/otelchi v0.5.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	github.com/tinylib/msgp v1.1.9 // indirect
	gitlab.dyninno.net/go-libraries/discovery-client v1.2.1 // indirect
	gitlab.dyninno.net/go-libraries/dreampass-client v1.1.2
	gitlab.dyninno.net/go-libraries/easycrypt v1.0.1 // indirect
	gitlab.dyninno.net/go-libraries/fluentdlogger v0.0.0-20230207094552-8a2a925b82b9 // indirect
	gitlab.dyninno.net/go-libraries/metrics v1.0.2
	go.opentelemetry.io/contrib v1.0.0 // indirect
	go.opentelemetry.io/otel v1.19.0 // indirect
	go.opentelemetry.io/otel/metric v1.19.0 // indirect
	go.opentelemetry.io/otel/sdk v1.19.0 // indirect
	go.opentelemetry.io/otel/trace v1.19.0 // indirect
	golang.org/x/net v0.14.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.12.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230711160842-782d3b101e98 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/driver/mysql v1.5.2 // indirect
	gorm.io/plugin/dbresolver v1.3.0 // indirect
	moul.io/http2curl v1.0.0
)
