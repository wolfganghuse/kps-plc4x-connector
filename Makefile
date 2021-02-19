SHELL := /usr/bin/env bash -o pipefail

# TODO: Replace with the actual connector name
# This controls the location of the cache.
PROJECT := kps-connector-go-template

# TODO: Replace with the actual connector name
export CONNECTOR_NAME := templateconnector

# TODO: Replace with the docker image tag
TAG=$(shell whoami)

# TODO: Replace with the docker repository URI
IMAGE_URI=registry.hub.docker.com/kps/connector/$(CONNECTOR_NAME):$(TAG)

### Everything below this line is meant to be static, i.e. only adjust the above variables. ###

UNAME_OS := $(shell uname -s)
UNAME_ARCH := $(shell uname -m)
# Buf will be cached to ~/.cache/buf-example.
CACHE_BASE := $(HOME)/.cache/$(PROJECT)
# This allows switching between i.e a Docker container and your local setup without overwriting.
CACHE := $(CACHE_BASE)/$(UNAME_OS)/$(UNAME_ARCH)
# The location where buf will be installed.
CACHE_BIN := $(CACHE)/bin
# Marker files are put into this directory to denote the current version of binaries that are installed.
CACHE_VERSIONS := $(CACHE)/versions
GO_FILES := $(shell \
	find . '(' -path '*/.*' -o -path './vendor' ')' -prune \
	-o -name '*.go' -print | cut -b3-)

# Update the $PATH so we can use buf directly
export PATH := $(abspath $(CACHE_BIN)):$(PATH)
# Update GOBIN to point to CACHE_BIN for source installations
export GOBIN := $(abspath $(CACHE_BIN))
# This is needed to allow versions to be added to Golang modules with go get
export GO111MODULE := on

GOLINT = $(GOBIN)/golint
GOSEC = $(GOBIN)/gosec

$(GOLINT):
	go get -u golang.org/x/lint/golint

$(GOSEC):
	go get -u github.com/securego/gosec/v2/cmd/gosec

.DEFAULT_GOAL := local

.PHONY: build
build:
	docker image build . -f Dockerfile -t $(IMAGE_URI) --build-arg connector_name=$(CONNECTOR_NAME)
	docker tag $(IMAGE_URI) baseimage

.PHONY: lint
lint: $(GOLINT) $(GOSEC)
	rm -rf lint.log
	echo "Checking formatting..."
	gofmt -d -s $(GO_FILES) 2>&1 | tee lint.log
	echo "Checking vet..."
	go vet ./... 2>&1 | tee -a lint.log
	echo "Checking lint..."
	$(GOLINT) $(GO_FILES) 2>&1 | tee -a lint.log
	echo "Checking security vulnerabilities..."
	$(GOSEC) -quiet ./... 2>&1 | tee -a lint.log

.PHONY: vendor
vendor:
	go mod vendor

.PHONY: test
test:
	go test -v ./...

.PHONY: cover
cover:
	go test -coverprofile=cover.out -coverpkg=./... ./...
	go tool cover -html=cover.out -o cover.html

# local is what we run when testing locally.
# This does breaking change detection against our local git repository.

.PHONY: local
local:
	make vendor
	make lint
	make build
	make test
	make cover

# clean deletes any files not checked in and the cache for all platforms.

.PHONY: clean
clean:
	git clean -xdf
	rm -rf $(CACHE_BASE)

.PHONY: publish
publish:
	docker push $(IMAGE_URI)
