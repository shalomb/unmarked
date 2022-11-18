#!/usr/bin/make -f

.ONESHELL:
SHELLFLAGS := -u nounset -ec

THIS_MAKEFILE := $(realpath $(lastword $(MAKEFILE_LIST)))
THIS_DIR      := $(shell dirname $(THIS_MAKEFILE))
THIS_PROJECT  := unmarked

VERSION := $(shell git rev-parse --short HEAD)
GITBRANCH := $(shell git branch --show-current)
BUILDTIME := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')
BUILDHOST := $(shell hostname -f)

GOLDFLAGS += -X main.Version=$(VERSION)
GOLDFLAGS += -X main.BuildTime=$(BUILDTIME)
GOLDFLAGS += -X main.Branch=$(GITBRANCH)
GOLDFLAGS += -X main.BuildHost=$(BUILDHOST)
GOFLAGS = -ldflags "$(GOLDFLAGS)"

.PHONY: serve watch

build: build-env build-darwin
	go build $(GOFLAGS)

build-darwin: build-env
	GOOS=darwin GOARCH=arm64 go build -o unmarked-darwin

build-env:
	go mod init unmarked
	go mod download

run: build
	./unmarked

serve:
	./unmarked

watch:
	watcher

version:
	./unmarked version
