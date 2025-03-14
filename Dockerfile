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

ENTRYPOINT ["/app/myapp"]