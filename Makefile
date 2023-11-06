SHELL := /bin/bash

.PHONY: all format lint test vet
.PHONY: install-go-test-coverage check-coverage

help:
	@echo "Please use \`make <target>\` where <target> is one of"
	@echo "  test                  to run all unit tests"
	@echo "  vet                   to do static check"
	@echo "  lint                  to run golangci lint"
	@echo "  format                to format code"

# only run unit tests, exclude e2e tests
test:
	go test -failfast $$(go list ./... | grep -v e2e) -covermode=atomic -coverprofile=./coverage.out -timeout 99999s
	# go test -cover ./...
	# go test -coverprofile=coverage.out ./...
	# go tool cover -html=coverage.out

vet:
	go vet ./...

lint:
	golangci-lint run --fix

format:
	gofmt -w -l .

check-coverage:
	@go-test-coverage --config=./.testcoverage.yml || true

install-go-test-coverage:
	go install github.com/vladopajic/go-test-coverage/v2@latest
