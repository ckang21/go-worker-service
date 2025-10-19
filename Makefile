.PHONY: run test build docker

run:
	go run ./cmd/server

test:
	go test ./...

build:
	go build -o bin/server ./cmd/server

docker:
	docker build -t go-worker-service:local .
