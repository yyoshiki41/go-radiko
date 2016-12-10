PKGS=$(shell go list ./... | grep -v examples)
BASE_FOLDERS=$(shell ls -d */ | grep -v vendor | grep -v testdata)

.PHONY: all help init test test-out

all: help

help:
	@echo "make init          #=> Run init scripts"
	@echo "make get-deps      #=> Install dependencies"
	@echo "make lint          #=> Verify tests"
	@echo "make lint          #=> Run golint"
	@echo "make vet           #=> Run go vet"
	@echo "make test          #=> Run tests"
	@echo "make test-out      #=> Run tests from outside Japan"

init: get-deps

test:
	go test $(PKGS)

test-out:
	go test -outjp $(PKGS)

test-ci:
	@echo "go test -outjp"
	@go test -outjp -race -coverprofile=coverage.txt -covermode=atomic $(PKGS)

verify: lint vet

lint:
	golint ./...

vet:
	go tool vet -all -structtags -shadow $(BASE_FOLDERS)

get-deps:
	@echo "go get go-radiko dependencies"
	@go get -v $(PKGS)
