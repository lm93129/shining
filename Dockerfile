FROM node:10.15 AS builder
WORKDIR /app
COPY . /app
RUN npm config set registry https://registry.npm.taobao.org && npm --verbose install
RUN npm run build:prod

FROM nginx:1.17.0-alpine-perl
RUN rm /etc/nginx/conf.d/default.conf
ADD default.conf /etc/nginx/conf.d/
COPY --from=builder /app/dist/ /usr/share/nginx/html/
