FROM golang:1.24.6

# Set the working directory inside the container to /app
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
# Download all Go module
RUN go mod download 
# Copy the rest of the application source code
COPY . .

RUN go build -o app .
EXPOSE 80
CMD ["./app"]