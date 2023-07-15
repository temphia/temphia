VERSION := $(shell git rev-parse --short HEAD)
BUILDTIME := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')
GOLDFLAGS += -X github.com/temphia/temphia/code/backend/xtypes.Version=$(VERSION)
GOLDFLAGS += -X github.com/temphia/temphia/code/backend/xtypes.Buildtime=$(BUILDTIME)
GOFLAGS = -ldflags "$(GOLDFLAGS)"

check: # check if binaries required are in $PATH
	@which wasm2wat
	@which docker
	@which docker-compose
	@which go
docker_run:
	cd ../cmd/dev && docker-compose up
server_run:
	cd .. && go run dev/dev.go
frontend_build:
	cd ../code/frontend && npm run build:@all
frontend_copy_files:
	cd ../code/frontend && npm run copy:build_files
backend_build:
	cd .. && go build -o build/temphia-cli $(GOFLAGS) .
build_all: frontend_build frontend_copy_files backend_build

.PHONY: check docker_run server_run frontend_build frontend_copy_files backend_build build_all install_cli backup_cli