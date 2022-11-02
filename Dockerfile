FROM node:10-alpine as builder
WORKDIR /app
COPY ./frontend /app
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
apk add make git && npm config set registry https://registry.npm.taobao.org && \
npm install && npm run build

FROM nginx:1.15-alpine
COPY --from=builder /app/dist/ /usr/share/nginx/html/
