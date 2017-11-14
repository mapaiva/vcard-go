install:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

test:
	make test-unit

test-unit:
	go test -v -cover --short ./...

lint:
	golint ./...

