install:
	go get

linter:
	golangci-lint run --timeout 2m0s

test-coverage:
	CGO_ENABLED=1 go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...