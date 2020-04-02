install:
	go get

linter:
	golangci-lint run --timeout=2m

test-coverage:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...