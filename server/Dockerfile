FROM golang:1.20 AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录：/build
WORKDIR /build

# 将代码复制到容器中
COPY . .

# 下载依赖信息
RUN go mod download && /bin/sh ./project_init.sh && go build -o server .

###################
# 接下来创建一个小镜像
###################
FROM alpine:latest

WORKDIR /go/src/base-system-backend/

COPY --from=builder /build/server /build/config.yaml /go/src/base-system-backend/

RUN apk update \
    && apk add tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && chmod +x /go/src/base-system-backend/server
