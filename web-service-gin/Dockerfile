# Stage 1: Build the Go binary
FROM golang:1.24.3 AS builder

WORKDIR /app

# Copy go.mod and go.sum, download deps
COPY src/go.mod src/go.sum ./
RUN go mod download

# Copy the rest of the code
COPY /src .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o album-api .

# Stage 2: Run the app with a minimal image
FROM alpine:latest

WORKDIR /root/
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/album-api .

EXPOSE 8080

CMD ["./album-api"]
