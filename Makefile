.PHONY: all test lint build run clean

all: lint test build

lint:
	@go vet ./...

test:
	@go test ./...

build:
	@mkdir -p ./bin
	@go build -o ./bin/demo ./stringer/cmd

run: build
	@./bin/demo

clean:
	@rm -rf ./bin
