FROM golang:1.10.3-alpine3.7 as builder
ENV PACKAGES go make git libc-dev bash
ENV REPO_PATH $GOPATH/src/github.com/irisnet/explorer/server
WORKDIR $REPO_PATH
COPY . $REPO_PATH/
RUN apk add --no-cache $PACKAGES && go get -u github.com/golang/dep/cmd/dep && make all

FROM alpine:3.7
ENV REPO_PATH /go/src/github.com/irisnet/explorer/server
COPY --from=builder $REPO_PATH/build/ /usr/local/bin/
