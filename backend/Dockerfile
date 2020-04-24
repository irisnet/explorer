FROM alpine:20190925 as builder
ENV PACKAGES go make git libc-dev bash
ENV GOPATH /root/go
ENV REPO_PATH $GOPATH/src/github.com/irisnet/explorer/backend
ENV GO111MODULE on
COPY . $REPO_PATH/
WORKDIR $REPO_PATH
RUN mkdir -p $REPO_PATH/build && apk add --no-cache $PACKAGES && make all

FROM alpine:3.7
ENV REPO_PATH /root/go/src/github.com/irisnet/explorer/backend
ENV TZ    Asia/Shanghai
WORKDIR /app/backend
RUN apk add -U tzdata && ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
COPY --from=builder $REPO_PATH/build/ /app/backend/
CMD ./irisplorer
