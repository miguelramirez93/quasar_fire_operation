# Compile stage
FROM golang:1.16 AS src
ENV APP_DIR=/go/src/github.com/miguelramirez93/quasar_fire_operation
ENV APP_NAME=quasar_fire_operation
RUN go get github.com/markbates/refresh
RUN go get -u github.com/swaggo/swag/cmd/swag
COPY . ${APP_DIR}
WORKDIR ${APP_DIR}
RUN go get ./...
RUN go test -cover ./...
RUN go test ./... -coverprofile cover.out
RUN go tool cover -func cover.out
RUN swag init

FROM src as dev
CMD ["refresh", "run"]

FROM src as build
RUN CGO_ENABLED=0 GOOS=linux go build
CMD ["./${APP_NAME}"]