FROM golang:alpine AS builder

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
RUN go mod download

# 将我们的代码编译成二进制可执行文件 bubble
RUN /bin/sh ./project_init.sh && go build -o server .

###################
# 接下来创建一个小镜像
###################
FROM alpine:latest

WORKDIR /go/src/base-system-backend/

COPY --from=builder /build/server /build/config.yaml /go/src/base-system-backend/

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime; \
    echo 'Asia/Shanghai' > /etc/timezone; \
    chmod +x /go/src/base-system-backend/server

# 需要运行的命令
# ENTRYPOINT ["/bubble", "conf/config.ini"]