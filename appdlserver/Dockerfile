# 构建镜像
FROM golang:1.14.1 AS builder
WORKDIR /app/build
COPY . /app/build
ENV GOPROXY=https://goproxy.cn,direct
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=jsoniter -o appdlserver

# 运行镜像
FROM alpine:3.11.2
WORKDIR /app
ENV PORT="3000"
ENV GIN_MODE="debug"
ENV LOG_LEVEL="debug"
COPY --from=builder /app/build/appdlserver /app/appdlserver
COPY ./docs /app/docs
ENTRYPOINT ["./appdlserver"]
LABEL Name=appdlserver Version=0.0.1
EXPOSE 3000
