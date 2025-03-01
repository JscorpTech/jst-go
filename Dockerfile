# Stage 1: Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Modul keshlash uchun go mod fayllarni alohida qo‘shamiz
COPY go.mod go.sum ./
RUN go mod download

# Qolgan kodni qo‘shamiz
COPY . .

# Dasturimizni kompilyatsiya qilamiz
RUN go build -o main ./cmd/main.go

# Stage 2: Runtime stage
FROM alpine:latest

WORKDIR /root/

# Oldingi stage'dan faqat kerakli binarini olib kelamiz
COPY --from=builder /app/main .
COPY .env .
COPY assets .

# Konteyner ishga tushganda bajariladigan buyruq
CMD ["./main"]
