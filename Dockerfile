# Stage 1: Build the Go application
FROM golang:1.21.6 AS builder

WORKDIR /go/src/app
COPY . .

RUN go run github.com/steebchen/prisma-client-go generate

RUN go build -o app cmd/http/main.go

# Stage 2: Run the application
FROM ubuntu:22.04

WORKDIR /app
COPY --from=builder /go/src/app/app /app
COPY ./docs/swagger.json /app/docs/swagger.json
COPY .env /app/.env
COPY serviceAccountKey.json /app/serviceAccountKey.json

CMD ["./app"]