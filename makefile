build_install_cli:
	go build -o build/temphia-cli cmd/cli/cli.go && cp build/temphia-cli ~/.bin