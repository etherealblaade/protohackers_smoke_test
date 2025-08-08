FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY smoke.go .

RUN go build -o tcpserver smoke.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/tcpserver .

EXPOSE 6359

CMD ["./tcpserver"]
