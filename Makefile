export GO111MODULE=on
SHELL:=/bin/bash

.PHONY: *

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/api ./cmd/
	chmod +x bin/api

up: build
	bin/api

tests: 
	go test -v -count=1 -covermode=count -coverprofile=coverage.out ./tests...

tests-unit: 
	go test -count=1 -covermode=count ./internal/...
	go test -count=1 -covermode=count ./http/...

tests-unit-internal:
	go test -v -count=1 -covermode=count -coverprofile=coverage.out ./internal/...
	go tool cover -html=coverage.out -o coverage.html
	open coverage.html >&- 2>&- || \
	xdg-open coverage.html >&- 2>&- || \
	gnome-open coverage.html >&- 2>&-

tests-unit-endpoints:
	go test -v -count=1 -covermode=count -coverprofile=coverage.out ./http/...
	go tool cover -html=coverage.out -o coverage.html
	open coverage.html >&- 2>&- || \
	xdg-open coverage.html >&- 2>&- || \
	gnome-open coverage.html >&- 2>&-

clean-tests:
	rm *.out