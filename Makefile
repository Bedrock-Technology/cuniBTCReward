.PHONY: generate build

generate:
	go generate ./api/openapi

build: generate
	go build ./...
