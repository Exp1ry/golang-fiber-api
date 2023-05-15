build:
	@go build -o bin/controller

run: build
	@./bin/controller

test:
	@go test -v ./...