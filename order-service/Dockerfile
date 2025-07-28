# Base image
FROM golang:1.24.5-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd
RUN go build -o order-service .

CMD ["./order-service"]