FROM registry.dyninno.net/dynatech/alpine-image/golang-1.21-dev AS development

ARG GITLAB_TOKEN=pat-token
RUN git config --global url."https://gitlab-ci-token:${GITLAB_TOKEN}@gitlab.dyninno.net/".insteadOf "https://gitlab.dyninno.net"
RUN go env -w GOPRIVATE=gitlab.dyninno.net
WORKDIR /home/app

COPY go.mod go.sum ./
RUN go mod download
RUN go install -mod=mod github.com/cosmtrek/air@v1.49.0
COPY . .

ENTRYPOINT ["air"]