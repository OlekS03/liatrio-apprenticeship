FROM golang:1.24.6

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .

RUN go build -o app .
EXPOSE 80
CMD ["./app"]