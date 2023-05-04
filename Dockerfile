FROM golang:1.17 as builder

#
RUN mkdir -p $GOPATH/src/gitlab.udevs.io/ucode/golang-app 
WORKDIR $GOPATH/src/gitlab.udevs.io/ucode/golang-app

# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    go mod vendor && \
    make build && \
    mv ./bin/golang-app /

FROM alpine
COPY --from=builder golang-app .
ENTRYPOINT ["/golang-app"]
