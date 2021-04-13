FROM golang:1.16.2-alpine

WORKDIR /app
COPY hello-world hello-world

EXPOSE 8080

COPY go.mod go.mod
COPY go.sum go.sum
COPY server.go server.go

RUN go build .

CMD ["/app/hello-world"]
