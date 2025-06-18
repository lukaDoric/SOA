FROM golang:alpine AS builder
WORKDIR /app
COPY . .
EXPOSE 8080
ENTRYPOINT ["go", "run", "main.go"]
