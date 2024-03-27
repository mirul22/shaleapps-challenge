# use official Golang image
FROM golang:1.16.3-alpine3.13

# maintainer
LABEL maintainer="Amirul Ikmal"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# set working directory
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

# Copy the source code
COPY . . 

# Download and install the dependencies
RUN go get -d -v ./...

# Build the Go app
RUN go build -o todo-api .

#EXPOSE the port
EXPOSE 8000

# Run the executable
CMD ["./todo-api"]