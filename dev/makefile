check: # check if binaries required are in $PATH
	@which wasm2wat
	@which cargo
	@which docker
	@which docker-compose
	@which go
	@which rustc

backend_docker_run:
	cd ../cmd/dev && docker-compose up
backend_server_run:
	cd .. && go run dev/dev.go
backend_server_tests_run:
	echo "hmm"

.PHONY: check backend_docker_run backend_server_run