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
RUN /bin/sh project_init.sh && go build -o base-system .

###################
# 接下来创建一个小镜像
###################
FROM debian:latest

WORKDIR /web/base-system-backend/

RUN mkdir -p /web/base-system-backend/initialize/internal/

COPY --from=builder /build/base-system /web/base-system-backend/
COPY --from=builder /build/config.yaml /web/base-system-backend/
COPY --from=builder /build/version.txt /web/base-system-backend/
COPY --from=builder /build/initialize/internal/privilege.json /web/base-system-backend/initialize/internal/

RUN set -eux; \
	apt-get update; \
	apt-get install -y netcat-openbsd; \
    chmox +x /web/base-system-backend/base-system; \
    /web/base-system-backend/base-system --system_init true

# 需要运行的命令
# ENTRYPOINT ["/bubble", "conf/config.ini"]