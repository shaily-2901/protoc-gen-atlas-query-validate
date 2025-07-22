GOPATH ?= $(HOME)/go
SRCPATH := $(patsubst %/,%,$(GOPATH))/src

# configuration for the protobuf gentool
PROJECT_ROOT := github.com/infobloxopen/protoc-gen-atlas-query-validate
SRCROOT_ON_HOST      := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
SRCROOT_IN_CONTAINER := /go/src/$(PROJECT_ROOT)
DOCKERPATH           := /go/src
DOCKER_RUNNER        := docker run --rm
DOCKER_RUNNER        += -v $(SRCROOT_ON_HOST):$(SRCROOT_IN_CONTAINER)
DOCKER_GENERATOR     := infoblox/atlas-gentool:latest
GENERATOR            := $(DOCKER_RUNNER) $(DOCKER_GENERATOR)

default: install

.PHONY: options
options:
	@$(GENERATOR)  \
		--go_out=":$(DOCKERPATH)" \
		options/query_validate.proto

.PHONY: install
install:
	go install

.PHONY: example
example: gentool
	$(DOCKER_RUNNER) infoblox/atlas-gentool:atlas-validate-query-dev \
	 --atlas-query-validate_out=":$(DOCKERPATH)" example/example.proto

test: example
	go test  ./...

.PHONY: vendor
vendor:
	go mod tidy && go mod vendor

.PHONY: gentool
gentool:
	docker build -t infoblox/atlas-gentool:atlas-validate-query-dev .
	docker image prune -f --filter label=stage=server-intermediate
