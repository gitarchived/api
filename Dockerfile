FROM golang:latest

COPY . .

RUN go build -o bin/api cmd/api/main.go

CMD ["./bin/api"]
