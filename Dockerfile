FROM golang:1.24.1-alpine AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o driven cmd/main.go

FROM debian:stable-slim
RUN apt-get update && apt-get install -y ca-certificates
COPY --from=builder /app/driven /usr/bin/driven
COPY static /static
CMD ["driven"]
