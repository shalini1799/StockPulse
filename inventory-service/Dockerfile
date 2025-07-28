# Use a minimal official Go image based on Alpine Linux
FROM golang:1.24.5-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy only go.mod and go.sum first to take advantage of Docker layer caching
# This avoids re-downloading all dependencies unless these files change
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download 

# Copy the rest of the application source code into the container
# The dot (.) means copy from the current context (host) into the working directory in container
COPY . .

# Change directory to /app/cmd (where the main.go or entry point is assumed to be)
WORKDIR /app/cmd

# Build the Go binary from the current directory (.)
# The dot refers to the package path to build (current directory)
RUN go build -o inventory-service .

# Run the compiled binary when the container starts
CMD ["./inventory-service"]