# Start from a base image with Go installed
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download and cache Go modules
RUN go mod download

# Copy the source code to the working directory
COPY . .

# Build the Go application
RUN go build -o server .

# Set environment variables for the PostgreSQL connection
ENV DB_HOST=postgres
ENV DB_PORT=5432
ENV DB_USER=yourusername
ENV DB_PASSWORD=yourpassword
ENV DB_NAME=yourdatabase

# Run the Go application
CMD ["./server"]
