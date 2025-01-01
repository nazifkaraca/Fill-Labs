# Base image
FROM golang:1.20

# Set the working directory
WORKDIR /app

# Copy everything into the container
COPY . .

# Download dependencies
RUN go mod tidy

# Build the application
RUN go build -o main .

# Expose the backend port
EXPOSE 8080

# Run the application
CMD ["./main"]
