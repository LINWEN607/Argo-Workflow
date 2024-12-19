FROM golang:1.22.5 AS builder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64
ENV GOPROXY "https://goproxy.cn,direct"
ENV GO111MODULE on

WORKDIR /app
COPY . .

RUN go mod tidy && go mod download

RUN go build -ldflags="-s -w" -o /home/app

FROM alpine:3.6

LABEL maintainer="Hydeli <hai_li@iot-dreamcatcher.com>"
# 切换软件源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
    && apk update \
    && apk add tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    apk clean

COPY --from=builder /home/app /home/app

EXPOSE 8080
ENTRYPOINT ["./home/app"]
