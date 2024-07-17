#build stage
FROM golang:1.22.4-alpine3.20 AS builder
WORKDIR /app
COPY . .
RUN sed -i 's/localhost/postgres/' /app/app.env

RUN go mod tidy
RUN go build -o . 
RUN apk --no-cache add curl 
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

## RUN STAGE
FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/* .

EXPOSE 8080
CMD ["/app/main"]

ENTRYPOINT ["/app/start.sh"]
