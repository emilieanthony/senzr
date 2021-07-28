include ../tools/golangci-lint/rules.mk

develop: run-database
	go run ./main.go

install:
	go mod tidy

test:
	go test ./...

lint: golangci-lint
