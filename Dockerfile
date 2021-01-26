FROM golang:1.15-alpine3.13
LABEL MAINTAINER = "bo"

ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.cn"
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64


# 创建目录,保存代码
RUN mkdir -p /opt/workspace/login
WORKDIR /opt/workspace/login
ADD go.mod .
ADD go.sum .

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o login .

EXPOSE 9060
CMD ["./login"]