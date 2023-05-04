# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.19 AS build-stage

WORKDIR /app

COPY go.mod ./

COPY *.go ./

RUN go build -o /golang-app


WORKDIR /

COPY --from=build-stage /golang-app /golang-app

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/golang-app"]