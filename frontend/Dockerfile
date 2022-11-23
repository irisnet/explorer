FROM node:10-alpine as builder
WORKDIR /app
COPY . /app
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
apk add make git && npm config set registry https://registry.npm.taobao.org && \
npm install && npm run build

FROM nginx:1.19-alpine
RUN echo -e 'server {\n\
  root /usr/share/nginx/html;\n\
  location / {\n\
    if ($request_filename ~* index.html|.*\.ico$)\n\
    {\n\
        add_header Cache-Control "no-cache";\n\
    }\n\
    try_files $URI $RUI/ /index.html;\n\
  }\n\
}' > /etc/nginx/conf.d/default.conf

COPY --from=builder /app/dist/ /usr/share/nginx/html/
CMD sh -c "exec nginx -g 'daemon off;'"
