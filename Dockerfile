<<<<<<< HEAD
FROM golang:1.21.6
=======
FROM golang:1.22.1-alpine3.18 AS builder
>>>>>>> main

RUN mkdir app

COPY . /app

WORKDIR /app

FROM alpine:latest

RUN go build -o main cmd/app/main.go

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app .

CMD ["/app/main"]