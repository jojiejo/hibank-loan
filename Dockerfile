FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o /hibank-loan

FROM alpine:latest

COPY --from=builder /hibank-loan /hibank-loan

COPY .env .env

EXPOSE 8080

CMD ["/hibank-loan"]