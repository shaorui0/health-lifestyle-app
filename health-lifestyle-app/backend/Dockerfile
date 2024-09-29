# Build Stage
FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main ./cmd/main.go

# Run Stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

# Set Environment Variables
ENV PORT=8000     DB_HOST=db     DB_PORT=5432     DB_USER=user     DB_PASSWORD=password     DB_NAME=health_app     JWT_SECRET=your_jwt_secret_key     OPENAI_API_KEY=your_openai_api_key_here

EXPOSE 8000

CMD ["./main"]
