FROM golang:alpine

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

ADD templates ./templates
COPY main.go .

RUN go build -o main .

EXPOSE 8080
CMD ["./main"]
