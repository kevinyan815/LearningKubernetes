# 第一阶段：构建阶段
FROM golang:alpine AS builder

# 安装必要的依赖，包括 Git 和 SSH 客户端
RUN apk update && apk add --no-cache git openssh

# 设置工作目录
WORKDIR /app

# 将私钥从构建参数传递到容器
ARG SSH_PRIVATE_KEY
RUN mkdir -p /root/.ssh && \
    echo "$SSH_PRIVATE_KEY" > /root/.ssh/id_rsa && \
    chmod 600 /root/.ssh/id_rsa

# 配置 Git 使用 SSH 连接
RUN mkdir -p /root/.ssh && \
    ssh-keyscan github.com >> /root/.ssh/known_hosts

# 配置环境变量
ENV GIT_SSH_COMMAND="ssh -i /root/.ssh/id_rsa -o StrictHostKeyChecking=no"

# 拉取代码并切换到指定的 Git tag
ARG GIT_TAG
RUN git clone git@github.com:go-study-lab/go-mall.git  .
RUN git checkout $GIT_TAG

# 安装 Go 项目的依赖
ENV GOPROXY=https://goproxy.cn,direct
RUN go mod tidy

# 构建 Go 项目
RUN go build -o go-mall .

# 第二阶段：运行阶段
FROM alpine:latest

# 安装运行时依赖
RUN apk --no-cache add ca-certificates tzdata
# 设置时区
ENV TZ=Asia/Shanghai

# 从构建阶段复制构建好的二进制文件
COPY --from=builder /app/go-mall /usr/local/bin/go-mall

# 设置容器启动时运行的命令
CMD ["go-mall"]