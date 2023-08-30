.PHONY: start

start: 
	TEMPHIA_DEV_PAGES_SERVER="http://localhost:5173/" go run ../cmd/dev/*.go

dev_build_frontend:
	cd ../code/frontend/ui && npm run portal:build
