FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

FROM alpine:3.20

# Tambahkan tzdata supaya timezone dikenali
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /root/

# Copy binary dari stage builder
COPY --from=builder /app/main .

# Copy environment file
COPY .env .env

# Set timezone container
ENV TZ=UTC

# Pastikan port sesuai docker-compose
EXPOSE 8080

CMD ["./main"]
