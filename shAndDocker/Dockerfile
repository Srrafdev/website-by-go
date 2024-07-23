# official Golang image based on Alpine Linux
FROM golang:1.22.3-alpine

# directory inside the container
WORKDIR /dockerize

# Copy all files from the current directory to the working directory inside the container
COPY . .

# Update the package list and install bash
RUN apk update && apk add bash

# Build the Go application
RUN go build -o main .

# Expose port 8080 to allow communication to/from the container on this port
EXPOSE 8080

# Add metadata to the image with the maintainer's name and email
LABEL name="rserraf"
LABEL email="srrafedev@gmail.com"

# run application
CMD ["./main"]
