FROM golang:1.23 AS builder

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main .
CMD ["./main"]