FROM golang:1.24.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-w -s" \
    -o driven \ 
    cmd/main.go

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /app/driven /usr/bin/driven
COPY static /static

CMD ["driven"]
