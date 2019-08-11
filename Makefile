SHELL = bash
PROJECT_ROOT := $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST)))))
GIT_COMMIT := $(shell git rev-parse HEAD)
GO_PKGS := $(shell go list ./...)

.PHONY: build
build:
	go build -o translate

.PHONY: test
test:
	go test -cover $(GO_PKGS)

.PHONY: clean
clean:
	@rm -f "$(PROJECT_ROOT)/translate"
