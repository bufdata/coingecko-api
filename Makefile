SHELL := /bin/bash

.PHONY: all format lint test vet

help:
	@echo "Please use \`make <target>\` where <target> is one of"
	@echo "  test                  to run all unit tests"
	@echo "  vet                   to do static check"
	@echo "  lint                  to run golangci lint"
	@echo "  format                to format code"

# only run unit tests, exclude e2e tests
test:
	go clean -testcache && go test -failfast $$(go list ./... | grep -v e2e) -timeout 99999s

vet:
	go vet ./...

lint:
	golangci-lint run --fix

format:
	gofmt -w -l .
