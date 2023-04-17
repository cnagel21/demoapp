
FROM golang:1.16

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build the Go app
RUN go build -o main .

# Expose port 8080 for the HTTP server
EXPOSE 8080

# Run the Go app
CMD ["./main"]
