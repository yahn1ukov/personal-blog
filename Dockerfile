FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum /app/

RUN go mod download

COPY . /app/

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /temp/api /app/cmd/main.go

FROM alpine:3.20

WORKDIR /app

COPY configs/config.yaml /app/config.yaml
COPY --from=builder /temp/api /app/api

EXPOSE 8000

CMD ["/app/api", "--config", "config.yaml"]
