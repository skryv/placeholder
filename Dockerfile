FROM golang:1.16.2-alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

EXPOSE 8080

WORKDIR /dist

COPY . .

RUN go get -v -t -d ./...
RUN go build .

FROM scratch

COPY --from=builder /dist/hello-world /

CMD ["/hello-world"]
