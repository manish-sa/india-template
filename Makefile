-include .env

.PHONY: oapi deps-lint deps-imports deps lint test imports migration migrate

DIR:=$(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST)))))
LOCAL_BIN:=$(DIR)/bin

help:
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-40s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Application commands

oapi: deps-oapi ## Generate oapi
	$(LOCAL_BIN)/oapi-codegen -package oapi -generate chi-server,types,strict-server -o ./internal/api/http/oapi/api.gen.go ./swagger/openapi.yaml
	$(LOCAL_BIN)/oapi-codegen -package oapi -generate spec -o ./internal/api/http/oapi/spec.gen.go ./swagger/openapi.yaml

OAPI_CODEGEN_VERSION ?= v1.14.0
deps-oapi:
ifeq ("$(wildcard $(LOCAL_BIN)/oapi-codegen)","")
	GOBIN=$(LOCAL_BIN) go install -mod=mod github.com/deepmap/oapi-codegen/cmd/oapi-codegen@$(OAPI_CODEGEN_VERSION)
endif

HUSKY_VERSION ?= v0.2.16
deps-husky:
ifeq ("$(wildcard $(LOCAL_BIN)/husky)","")
	GOBIN=$(LOCAL_BIN) go install -mod=mod github.com/automation-co/husky@$(HUSKY_VERSION)
endif

GOLANG_CI_LINT_VERSION ?= v1.54.2
deps-lint:
ifeq ("$(wildcard $(LOCAL_BIN)/golangci-lint)","")
	GOBIN=$(LOCAL_BIN) go install -mod=mod github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANG_CI_LINT_VERSION)
endif

IMPORTS_REVISER_VERSION ?= v3.4.3
deps-imports:
ifeq ("$(wildcard $(LOCAL_BIN)/goimports-reviser)","")
	GOBIN=$(LOCAL_BIN) go install -mod=mod github.com/incu6us/goimports-reviser/v3@$(IMPORTS_REVISER_VERSION)
endif

MIGRATE_VERSION ?= v4.16.2
deps-migrate:
ifeq ("$(wildcard $(LOCAL_BIN)/migrate)","")
	GOBIN=$(LOCAL_BIN) go install -tags $(DB_DRIVER) -mod=mod github.com/golang-migrate/migrate/v4/cmd/migrate@$(MIGRATE_VERSION)
endif

AIR_VERSION ?= v1.45.0
deps-air:
ifeq ("$(wildcard $(LOCAL_BIN)/air)", "")
	GOBIN=$(LOCAL_BIN) go install -mod=mod github.com/cosmtrek/air@$(AIR_VERSION)
endif

deps: deps-lint deps-imports deps-migrate deps-oapi deps-air husky ## Deps
	go mod tidy

generate: oapi
	go generate ./...

export PROJECT_NAME ?= lbc
PROJECT_PATH ?= gitlab.dyninno.net/trevolution/ancillaries/lbp/lbc-service
LD_FLAGS ?= "-X $(PROJECT_PATH)/internal/info.serviceName=$(PROJECT_NAME)"
build: generate
	go build -trimpath -ldflags $(LD_FLAGS) -buildvcs=false -v -o $(LOCAL_BIN)/app ./cmd

watch:
	$(LOCAL_BIN)/air api

lint: deps-lint ## Lint
	GOFLAGS=-buildvcs=false $(LOCAL_BIN)/golangci-lint run --timeout 10m

husky: deps-husky ## Lint
	$(LOCAL_BIN)/husky install

COVERAGE_EXCLUDE ?= 'config\|gen.go'
COVERAGE_PATHS ?= './...'
RUN_COUNT ?= 5
test: ## Unit tests
	go clean -testcache
	go test $(COVERAGE_PATHS) -coverprofile=coverage.raw -covermode=atomic -coverpkg=$(strip$(COVER_PKG)) -count=$(RUN_COUNT) -race
	grep -v $(COVERAGE_EXCLUDE) coverage.raw  > coverage && rm coverage.raw
	go tool cover -func coverage
	grep -sqFx "coverage" .gitignore || echo "coverage" >> .gitignore

imports: deps-imports ## Imports
	@find . \
		\( -type d -path './vendor' -o -type d -path './.dev' -o -type d -path './.git' -o -type d -path './internal/api/http/oapi' \) -prune -o \
		-type f \
		-name \*.go \
	-exec sh -c 'echo "Processing: {}"; $(LOCAL_BIN)/goimports-reviser -rm-unused -set-alias -format -company-prefixes gitlab.dyninno.net/ "{}"' \;

MIGRATION_NAME ?= migration
MIGRATIONS_DIR ?= migrations
MIGRATIONS_EXT ?= sql

migration: deps-migrate ## Migration create
	$(LOCAL_BIN)/migrate create -ext $(MIGRATIONS_EXT) -dir $(MIGRATIONS_DIR) $(MIGRATION_NAME)

MIGRATE_ARGS ?= up
MIGRATE_URL ?= $(DB_DRIVER)://$(DB_USER):$(DB_PASS)@tcp($(DB_MASTER_HOST):$(DB_MASTER_PORT))/$(DB_NAME)
.migrate: deps-migrate
	$(LOCAL_BIN)/migrate -path $(MIGRATIONS_DIR) -verbose -database "$(MIGRATE_URL)" $(MIGRATE_ARGS)

migrate: ## Migration run
	/bin/sh ./scripts/migrate.sh

up: ## docker up
	docker compose up -d

down: ## docker down
	docker-compose down --remove-orphans
