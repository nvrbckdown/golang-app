# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.19-alpine AS build-stage

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping
# Deploy the application binary into a lean image
FROM alpine AS build-release-stage

WORKDIR /

COPY --from=build-stage /docker-gs-ping /docker-gs-ping
COPY --from=build-stage /app/temp /temp
EXPOSE 8080

ENTRYPOINT ["/docker-gs-ping"]
