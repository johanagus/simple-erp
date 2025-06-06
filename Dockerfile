# Start from the official Golang base image
FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o simple-erp ./cmd/main.go

EXPOSE 8080

CMD ["./simple-erp"]
