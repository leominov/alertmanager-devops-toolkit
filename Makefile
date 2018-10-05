.PHONY: release lint

release:
	@./.release.sh

lint:
	@go run *.go --render > alertmanager.yml
	@cat -n alertmanager.yml
	@echo "Result:"
	@go run *.go --lint
