FROM golang:1.22.1-alpine3.19

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o main
EXPOSE 3000
CMD ["./main"]