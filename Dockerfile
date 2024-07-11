FROM golang:latest

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64
ENV GOPROXY "https://goproxy.cn,direct"
ENV GO111MODULE on

WORKDIR /app
COPY . .

RUN go build -ldflags="-s -w" -o /home/app

EXPOSE 8000
ENTRYPOINT ["./go-gin-example"]
