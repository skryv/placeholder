FROM golang:1.16.2

WORKDIR /app
COPY hello-world hello-world

EXPOSE 8080

CMD ["/app/hello-world"]
