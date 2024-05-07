FROM golang:1.21.6

RUN mkdir app

COPY . /app

WORKDIR /app

FROM alpine:latest

RUN go build -o main cmd/app/main.go

CMD ["/app/main"]

EXPOSE 8080