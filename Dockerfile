FROM golang:1.16.2-alpine

WORKDIR /app
COPY hello-world hello-world

EXPOSE 8080

CMD ["/app/hello-world"]
