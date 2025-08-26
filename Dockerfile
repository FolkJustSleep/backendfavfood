# Use the official Go image as the base image
FROM golang:1.24-alpine AS builder

# Set the working directory in the container
WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go mod download

# build binary แบบ static
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Final image
FROM alpine:latest

WORKDIR /app

# เอา binary ที่ build แล้วมาใส่
COPY --from=builder /app/main .

# Expose port 8080
EXPOSE 1818

# Define the entry point for the container
CMD ["./main"]
