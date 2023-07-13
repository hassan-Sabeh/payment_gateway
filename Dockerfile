# Use the official Go image as the base image
FROM golang:1.18

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Download the missing testify package
# RUN go mod download github.com/stretchr/testify

# Copy the project source code to the working directory
COPY . .

# Build the Go application
RUN go build -o main ./cmd/service

# Set the command to run the Go application
CMD ["./main"]