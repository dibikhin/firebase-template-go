.PHONY: all test build run clean

local_server = local_server

all: clean test build

test:
	@echo "\nTesting..."
	go test -v ./...

build:
	@echo "\nBuilding..."
	@go version
	go build -o ${local_server} ./cmd/https

run: ${local_server}
	@echo "\nRunning..."
	./${local_server}

clean:
	@echo "\nCleaning up..."
	rm -f ${local_server}
