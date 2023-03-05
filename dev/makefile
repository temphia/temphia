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
	cd .. && go build -o build/temphia-cli cmd/cli/cli.go && cp build/temphia-cli ~/.bin
backup_cli:
	cd .. && go build -o build/temphia-cli cmd/cli/cli.go && cp build/temphia-cli ~/Desktop/temphia_run

.PHONY: check docker_run server_run build_install_cli backup_cli