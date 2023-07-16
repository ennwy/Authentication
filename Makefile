BIN := "./bin/blog"

build:
	go build -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/...

run: build
	HTTP_HOST="localhost" HTTP_PORT="8080" $(BIN)

test:
	go test -v ./...

up:
	sudo docker-compose -f ./deployments/docker-compose.yaml up --build

down:
	sudo docker-compose -f ./deployments/docker-compose.yaml down

.PHONY: build run test up down