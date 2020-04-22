.PHONY: all test build run clean

all: test build

test:
	go test -v ./...

build:
	@echo "\nBuilding..."
	go build -o local_server ./cmd/https

run:
	@echo "\nRunning..."
	./local_server

clean: local_server
	@echo "Cleaning up..."
	rm local_server

