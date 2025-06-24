FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

FROM alpine:latest

# Install curl and ca-certificates for health checks
RUN apk add --no-cache curl ca-certificates

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs
COPY --from=builder /app/db/migrations ./db/migrations

EXPOSE 8080

ENTRYPOINT ["./main"]
CMD ["--db-host=db", "--db-port=5432", "--db-user=hear", "--db-password=hear", "--db-name=hear"] 