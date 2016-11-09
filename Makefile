PKGS=$(shell go list ./...)

.PHONY: all help init test test-out

all: help

help:
	@echo "make init          #=> Run init scripts"
	@echo "make test          #=> Run tests"
	@echo "make test-out      #=> Run tests from outside Japan"

init: get-deps

test:
	go test $(PKGS)

test-out:
	go test -outjp $(PKGS)

test-ci:
	@echo "go test"
	@go test -outjp -race -coverprofile=coverage.txt -covermode=atomic $(PKGS)

get-deps:
	@echo "go get go-radiko dependencies"
	@go get -v $(PKGS)
