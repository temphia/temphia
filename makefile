.PHONY: clean changelog snapshot
.PHONY: deps

# Check for required command tools to build or stop immediately
EXECUTABLES = git go find pwd
K := $(foreach exec,$(EXECUTABLES),\
        $(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH")))

VERSION ?= $(shell git describe --tags `git rev-list --tags --max-count=1`)
BINARY = temphia

BUILDDIR = ../build
GITREV = $(shell git rev-parse --short HEAD)
BUILDTIME = $(shell date +'%FT%TZ%z')
GOLDFLAGS += -X github.com/temphia/temphia/code/backend/xtypes.Version=$(VERSION)
GOLDFLAGS += -X github.com/temphia/temphia/code/backend/xtypes.Buildtime=$(BUILDTIME)
GOFLAGS = -ldflags "$(GOLDFLAGS)"
GO_BUILDER_VERSION=latest

deps:
	go get -u github.com/git-chglog/git-chglog/cmd/git-chglog
	go get -u golang.org/x/tools/cmd/goimports

clean:
	rm -rf $(shell pwd)/$(BUILDDIR)/

changelog:
	git-chglog $(VERSION) > CHANGELOG.md

snapshot:
	docker run --rm --privileged \
		-v $(CURDIR):/temphia \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v $(GOPATH)/src:/go/src \
		-w /temphia \
		ghcr.io/gythialy/golang-cross:$(GO_BUILDER_VERSION) --snapshot --clean --rm-dist
		