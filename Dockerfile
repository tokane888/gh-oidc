# Build stage
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/app

# Run stage
FROM gcr.io/distroless/base-debian11

WORKDIR /app

COPY --from=builder /app/app .
COPY .env .env

# TODO: port周りの扱いは追って検討
# EXPOSE 8080

ENTRYPOINT ["./app"]
