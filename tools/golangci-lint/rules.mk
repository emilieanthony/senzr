path := ../tools/golangci-lint
arch := $(shell uname -m)
os := $(shell uname -s)
version := 1.41.1

ifeq ($(shell uname),Linux)
arch = 386
os = linux
else ifeq ($(shell uname),Darwin)
os = darwin
else
$(error unsupported OS: $(shell uname))
endif

.PHONY: golangci-lint
golangci-lint:
	curl -L https://github.com/golangci/golangci-lint/releases/download/v$(version)/golangci-lint-$(version)-$(os)-$(arch).tar.gz | tar xvzf - -C $(path)/
	chmod 755 $(path)/golangci-lint-$(version)-$(os)-$(arch)/golangci-lint
	./$(path)/golangci-lint-$(version)-$(os)-$(arch)/golangci-lint run --timeout 5m
