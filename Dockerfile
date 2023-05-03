# workspace (GOPATH) configured at /go
FROM golang:1.20.0 as builder

#
RUN mkdir -p $GOPATH/src/golang-app
WORKDIR $GOPATH/src/golang-app

# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN go build -o golang-app

FROM alpine
COPY --from=builder golang-app .
ENTRYPOINT ["/golang-app"]