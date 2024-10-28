FROM golang:1.22.3-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o main main.go

FROM alpine:latest

COPY --from=builder /app/main /app/main

ENTRYPOINT ["./main"]
