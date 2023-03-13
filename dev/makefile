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

build_install_cli:
	cd .. && go build -o build/temphia-cli $(GOFLAGS) cmd/cli/cli.go && cp build/temphia-cli ~/.bin
backup_cli:
	cd .. && go build -o build/temphia-cli cmd/cli/cli.go && cp build/temphia-cli ~/Desktop/temphia_run

.PHONY: check docker_run server_run build_install_cli backup_cli