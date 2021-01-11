install:
	go get

linter:
	golangci-lint run

test-coverage:
	CGO_ENABLED=1 go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...