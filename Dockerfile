FROM golang:alpine

WORKDIR "/app"

COPY ./ ./

CMD ["go", "run", "docker-go-http"]