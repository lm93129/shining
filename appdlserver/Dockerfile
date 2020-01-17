#build stage
FROM golang:1.13.5 AS builder
WORKDIR /app/build
COPY . /app/build
RUN apk add --no-cache git
RUN go install -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o appdlserver

FROM registry.cn-hangzhou.aliyuncs.com/lm93129/alpine:3.11.2

WORKDIR /app
COPY --from=builder /app/build/appdlserver /app/appdlserver
ENTRYPOINT ["./appdlserver"]
LABEL Name=appdlserver Version=0.0.1
EXPOSE 3000
