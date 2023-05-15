# Use an official Golang runtime as a parent image
FROM --platform=amd64 golang:1.20-alpine

WORKDIR /app

# Set the permissions for the working directory
RUN chown -R nobody:nobody /app
RUN chmod -R 777 /app

# Copy the current directory contents into the container at /app
COPY . .

# Build the Go app
RUN go build -o app

# Set execute permissions for the binary
RUN chmod +x ./app

# Expose port 3000
EXPOSE 8080

# Run the Go app
CMD ["./app"]
