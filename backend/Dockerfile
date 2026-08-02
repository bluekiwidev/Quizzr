FROM golang:tip-alpine AS builder
WORKDIR /app

COPY src/ /app/

RUN go mod tidy
RUN go build -o main .

FROM alpine:latest

ENV MYSQL_HOST=localhost
ENV MYSQL_PORT=3306
ENV MYSQL_USER=user
ENV MYSQL_PASSWORD=
ENV MYSQL_DATABASE=database

COPY --from=builder /app/main /app/main

CMD ["/app/main"]