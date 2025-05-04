FROM golang:1.24 as builder

WORKDIR /app

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the Go Fiber application with static linking
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/app

FROM alpine:3.9.6

# Install certificates and timezone data
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Copy the built binary from the builder stage and copy environment variables
COPY --from=builder /app/main .
COPY .env .

EXPOSE 8080

CMD ["./main"]
