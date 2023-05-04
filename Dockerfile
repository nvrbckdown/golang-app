# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.19 AS build-stage

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /golang-app


WORKDIR /

COPY --from=build-stage /golang-app /golang-app

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/golang-app"]