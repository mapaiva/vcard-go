test:
	make test-unit

test-unit:
	go test -v -cover --short ./...

lint:
	golint ./...

