check: # check if binaries required are in $PATH
	@which wasm2wat
	@which docker
	@which docker-compose
	@which go
docker_run:
	cd ../cmd/dev && docker-compose up
server_run:
	cd .. && go run dev/dev.go

.PHONY: check docker_run server_run