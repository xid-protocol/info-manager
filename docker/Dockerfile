# 多阶段构建
FROM golang:1.24-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制 go mod 文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

# 运行阶段
FROM alpine:latest

# 安装必要的系统依赖
RUN apk --no-cache add ca-certificates tzdata

# 设置时区
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

WORKDIR /root/

# 从构建阶段复制可执行文件
COPY --from=builder /app/main .

# 创建配置目录
RUN mkdir -p /opt/xidp/conf
RUN mkdir -p /opt/xidp/logs

# 暴露端口
EXPOSE 8080

# 运行应用
CMD ["./main", "-c", "/opt/xidp/conf/config.yml"] 