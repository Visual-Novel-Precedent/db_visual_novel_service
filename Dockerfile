FROM golang:1.23-bookworm AS builder

WORKDIR /build

# Копируем только необходимые файлы для сборки
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь исходный код
COPY . .

# Сборка с оптимизациями
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main ./cmd/service/main

FROM scratch AS production

WORKDIR /app

# Копируем только собранный бинарник
COPY --from=builder /build/main .

EXPOSE 8080

CMD ["./main"]