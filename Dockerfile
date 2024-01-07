# Use the official Golang image as the base image
FROM golang:1.19

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application files into the container
COPY app/ ./

# Build the Go application
RUN go build main.go

# Expose the port the application runs on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
