# Build stage
FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o todo-server

# Run stage
FROM alpine:latest

RUN addgroup -S todo-api && adduser -S todo-api -G todo-api

WORKDIR /home/todo-api/app

COPY --from=builder /app/todo-server /app/.env .

RUN chown -R todo-api:todo-api /home/todo-api

USER todo-api

EXPOSE 8080

CMD ["./todo-server"]
