PKGS=$(shell go list ./... | grep -v examples)

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
	@echo "" > coverage.txt; \
	for d in $(PKGS); do \
		go test -outjp -coverprofile=profile.out -covermode=atomic $$d || exit 1; \
		[ -f profile.out ] && cat profile.out >> coverage.txt && rm profile.out || true; \
	done

get-deps:
	@echo "go get go-radiko dependencies"
	@go get -v $(PKGS)
