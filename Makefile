install:
	go mod tidy
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.1

test:
	make test-unit

test-unit:
	go test -v -cover --short ./...

lint:
	golangci-lint run

