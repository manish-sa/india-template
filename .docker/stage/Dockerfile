FROM registry.dyninno.net/dynatech/alpine-image/golang-1.21-runner-travel AS builder

ARG PIPELINE_TOKEN
RUN git config --global url."https://gitlab-ci-token:${PIPELINE_TOKEN}@gitlab.dyninno.net/".insteadOf "https://gitlab.dyninno.net"
RUN go env -w GOPRIVATE=gitlab.dyninno.net

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY .docker/stage/.env ./.env
RUN make deps-migrate

ARG CI_PROJECT_ID
ARG CI_PROJECT_PATH
ARG CI_PROJECT_NAME
ARG CI_PROJECT_NAMESPACE
ARG CI_COMMIT_REF_SLUG
ARG CI_SERVER_HOST
ARG CI_COMMIT_REF_NAME

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags " \
    -X '$CI_SERVER_HOST/$CI_PROJECT_PATH/internal/info.serviceName=$CI_PROJECT_NAME'  \
    -X '$CI_SERVER_HOST/$CI_PROJECT_PATH/internal/info.namespace=$CI_PROJECT_NAMESPACE'  \
    -X '$CI_SERVER_HOST/$CI_PROJECT_PATH/internal/info.version=$CI_COMMIT_REF_SLUG'  \
    -X '$CI_SERVER_HOST/$CI_PROJECT_PATH/internal/info.gitlabProjectID=$CI_PROJECT_ID'  \
    -X '$CI_SERVER_HOST/$CI_PROJECT_PATH/internal/info.buildDate=$(date -u +%FT%TZ)' \
    -X '$CI_SERVER_HOST/$CI_PROJECT_PATH/internal/info.gitLog=$(git log --decorate --oneline -n5 2> /dev/null | base64 | tr -d '\n')' \
    -X '$CI_SERVER_HOST/$CI_PROJECT_PATH/internal/info.gitHash=$(git log --format="%h" -n 1 2> /dev/null)' \
    -X '$CI_SERVER_HOST/$CI_PROJECT_PATH/internal/info.gitBranch=$CI_COMMIT_REF_NAME'" \
    -o /app ./cmd

FROM registry.dyninno.net/dynatech/alpine-image/alpine-3/alpine:3.18.4

RUN apk --no-cache add curl make

COPY .docker/stage/.env ./.env
COPY --from=builder app /app/app
COPY --from=builder /build/bin/. bin/.
COPY --from=builder /build/Makefile Makefile
COPY --from=builder /build/swagger swagger
COPY --from=builder /build/scripts scripts
COPY --from=builder /build/migrations migrations

ENTRYPOINT ["/app/app", "api"]
