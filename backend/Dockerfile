FROM alpine:20190925 as builder
ENV PACKAGES go make git libc-dev bash
ENV REPO_PATH $GOPATH/src/github.com/irisnet/explorer/server
ENV GO111MODULE on
WORKDIR $REPO_PATH
COPY . $REPO_PATH/
RUN apk add --no-cache $PACKAGES && make all

FROM alpine:3.7
ENV REPO_PATH /go/src/github.com/irisnet/explorer/server
ENV TZ           Asia/Shanghai
COPY --from=builder $REPO_PATH/build/ /usr/local/bin/
