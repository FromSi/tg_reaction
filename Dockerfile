FROM golang:1.24-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp ./cmd/tg_bot/main.go

FROM alpine:edge

WORKDIR /app

COPY --from=build /app/myapp .
COPY --from=build /app/config.json .

HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 \
    CMD pgrep -f "/app/myapp" > /dev/null || exit 1

ENTRYPOINT ["/app/myapp"]