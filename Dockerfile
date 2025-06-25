FROM golang:tip-alpine3.21

WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./
RUN apk update && apk upgrade --no-cache && go mod download

# Copy seluruh source code ke dalam container
COPY . .

# Build dari folder cmd
RUN go build -o simple-erp ./cmd

EXPOSE 8000

CMD ["./simple-erp"]