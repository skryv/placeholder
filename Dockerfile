FROM golang:1.24.1-alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

EXPOSE 8080

WORKDIR /dist

COPY . .

RUN go get -v -t -d ./...
RUN go build .

FROM scratch

COPY --from=builder /dist/placeholder /

CMD ["/placeholder"]
