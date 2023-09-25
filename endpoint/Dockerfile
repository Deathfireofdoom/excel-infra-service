# Use the official Go image
FROM golang:1.21

workdir /app

# Copy go.mod and go.sum files
COPY ./src/go.mod ./src/go.sum ./

# Download all dependencies
RUN go mod download

# Copy the entire source code
COPY ./src .

# Build the application
RUN CGO_ENABLED=1 go build -o main ./cmd/main.go

# Run the application
CMD ["./main"]