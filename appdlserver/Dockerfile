# 构建镜像
FROM dockerhub.azk8s.cn/library/golang:1.13.5 AS builder
WORKDIR /app/build
COPY . /app/build
ENV GOPROXY=https://goproxy.cn,direct
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o appdlserver

# 运行镜像
FROM dockerhub.azk8s.cn/library/alpine:3.11.2
WORKDIR /app
ENV PROT="3000"
ENV GIN_MODE="debug"
ENV LOG_LEVEL="debug"
COPY --from=builder /app/build/appdlserver /app/appdlserver
COPY ./docs /app/docs
ENTRYPOINT ["./appdlserver"]
LABEL Name=appdlserver Version=0.0.1
EXPOSE 3000
