FROM golang:1.18-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download

RUN go build -o blog-backend cmd/main.go

EXPOSE 3500

CMD ["./blog-backend"]