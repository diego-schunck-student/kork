FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY main.go .
# Inicializa o módulo e baixa as dependências do Prometheus no build
RUN go mod init http-server-projeto-korp && \
    go get github.com/prometheus/client_golang/prometheus && \
    go get github.com/prometheus/client_golang/prometheus/promauto && \
    go get github.com/prometheus/client_golang/prometheus/promhttp && \
    go build -o server main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
