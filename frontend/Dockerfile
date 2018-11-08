FROM node:10-alpine as builder
WORKDIR /app
COPY . /app
RUN npm i yarn -g && yarn install && yarn build

FROM nginx:1.15-alpine
COPY --from=builder /app/dist/ /usr/share/nginx/html/
