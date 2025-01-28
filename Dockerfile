# Use official Golang image
FROM golang:1.20

# Set working directory
WORKDIR /app

# Copy the source code
COPY . .

# Install dependencies
RUN go mod tidy

# Build the Go application
RUN go build -o main .

# Command to run the executable
CMD ["/app/main"]
