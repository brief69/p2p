# Use the official Golang image from the Docker Hub
FROM golang:1.16-alpine

# Install the git package which will be used to fetch the dependencies of the project
RUN apk add --no-cache git

# Create a directory inside the container to store all our application and then make it the working directory.
RUN mkdir /app
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
