build:
	golangci-lint run
	golangci-lint fmt
	go build ./...
	go test ./...

test:
	go test ./...

lint:
	golangci-lint run

format:
	golangci-lint fmt

	