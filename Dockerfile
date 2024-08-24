# Stage 1: Build the Go application
FROM golang:1.21.6 AS builder

WORKDIR /go/src/app
COPY . .

RUN go build -o app cmd/http/main.go

# Stage 2: Run the application
FROM ubuntu:22.04

WORKDIR /app
COPY --from=builder /go/src/app/app /app
COPY .env /app/.env

CMD ["./app"]