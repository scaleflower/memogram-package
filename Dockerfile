# Build stage
FROM cgr.dev/chainguard/go:latest AS builder
WORKDIR /app

# 拷贝所有文件（顺序很重要，先复制才能 tidy）
COPY . .

# 整理依赖
RUN go mod tidy

# 编译可执行文件（路径为 bin/memogram）
RUN CGO_ENABLED=0 go build -o memogram ./bin/memogram

# Final stage
FROM cgr.dev/chainguard/static:latest-glibc
WORKDIR /app

COPY --from=builder /app/memogram .
COPY .env.example .env

ENTRYPOINT ["/app/memogram"]
