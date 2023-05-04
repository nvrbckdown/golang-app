# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.19 AS build-stage

WORKDIR /app

COPY . ./

RUN go build -o /golang-app

FROM alpine
COPY --from=build-stage /golang-app /golang-app

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/golang-app"]
