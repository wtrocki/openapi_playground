# Requires golangci-lint to be installed @ $(go env GOPATH)/bin/golangci-lint
# https://golangci-lint.run/usage/install/

lint:
	golangci-lint run ./...
.PHONY: lint

generate:
	./scripts/generate.sh 
.PHONY: generate

mock:
	npx @stoplight/prism-cli mock .openapi/content-type.json
.PHONY: generate

example:
	go run ./examples/contenttype.go
.PHONY: example

help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
.PHONY: help

