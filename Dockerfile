FROM golang:1.24.1

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the Go app
RUN go build -o main .

# Command to run the app
CMD ["./main","server"]