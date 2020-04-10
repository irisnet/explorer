FROM alpine:20190925 as builder
ENV PACKAGES go make git libc-dev bash
ENV REPO_PATH $GOPATH/src/github.com/irisnet/explorer/server
ENV GO111MODULE on
WORKDIR $REPO_PATH
COPY . $REPO_PATH/
RUN apk add --no-cache $PACKAGES && apk add -U tzdata && make all
COPY /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN apk del tzdata

FROM alpine:3.7
ENV REPO_PATH /go/src/github.com/irisnet/explorer/server
COPY --from=builder $REPO_PATH/build/ /usr/local/bin/
