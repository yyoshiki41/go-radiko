PKG=$(shell go list ./...)

.PHONY: all help test test-out

all: help

help:
	@echo "make test          #=> Run tests"
	@echo "make test-out      #=> Run tests from outside Japan"

test:
	go test $(PKG)

test-out:
	go test -outjp $(PKG)

test-ci:
	go test -outjp -race -coverprofile=coverage.txt -covermode=atomic $(PKG)
