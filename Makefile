install:
	go get ./...
	go install github.com/envoyproxy/protoc-gen-validate@v0.6.7

linter:
	golangci-lint run --timeout 2m0s

test-coverage:
	CGO_ENABLED=1 go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

.PHONY: proto
proto:
	mkdir -p proto/basic
	protoc \
		-I proto \
		-I ${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.6.7 \
		--go_out=proto/basic \
        --go-grpc_out=proto/basic \
        --validate_out="lang=go:proto/basic" \
        proto/*.proto