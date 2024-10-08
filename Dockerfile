FROM golang:1.21-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

FROM alpine:edge

WORKDIR /app

COPY --from=build /app/myapp .

ENTRYPOINT ["/app/myapp"]