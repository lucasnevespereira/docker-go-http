# Docker-go-http

Sample Docker Image that starts a Golang HTTP server.

- Inited a go module and a main.go with an http server.

- Created a Dockerfile from golang alpine linux, that copied my go files to the workdir and then runs the go server.

### Build the container from docker image

docker build -t lucasnpereira/docker-go-http:latest .

# Run container

docker run -p 4001:4000 lucasnpereira/docker-go-http

# Visit https://localhost:4001
