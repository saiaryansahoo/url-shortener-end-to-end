# Build the backend
FROM golang:1.23 AS builder
WORKDIR /app/backend  
COPY backend/ .        
RUN go mod tidy && go build -o /app/url-shortener .

# Final image
FROM alpine:latest
WORKDIR /root/
RUN apk add --no-cache libc6-compat  # Add missing dependencies
COPY --from=builder /app/url-shortener .
COPY frontend/ ./frontend
EXPOSE 8080
CMD ["./url-shortener"]

