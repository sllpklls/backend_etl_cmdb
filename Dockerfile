FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main .

# Final stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]
