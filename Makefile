generate:
	@go generate ./...
.PHONY: generate

install:
	@vgo mod vendor

build: generate
	@echo "====> Build telnetor cli"
	@go build -o ./bin/badge-service main.go
.PHONY: build

release:
	@echo "====> Build and release"
	@go get github.com/goreleaser/goreleaser
	@goreleaser
.PHONY: release

test:
	@go test ./...
.PHONY: test

test.cov:
	@go test ./... -coverprofile=coverage.txt -covermode=atomic
.PHONY: test.cov
