ARG GO_VERSION=1.22.5
ARG GO_BASE_IMAGE=alpine

FROM golang:${GO_VERSION}-${GO_BASE_IMAGE} AS builder
WORKDIR /app
COPY ./backend .
RUN go mod download
<<<<<<< HEAD
RUN go build -o bin/calculator ./cmd/calculator
RUN go build -o bin/migrator ./cmd/migrator

FROM alpine:latest
WORKDIR /app
COPY --from=builder ./app/bin ./bin
COPY --from=builder ./app/internal/data ./data
=======
RUN go build -o bin/main ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder ./app/bin ./bin
>>>>>>> master
