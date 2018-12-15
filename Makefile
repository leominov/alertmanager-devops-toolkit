.PHONY: release lint

release:
	@./.release.sh

lint:
	@SLACK_API_URL=http://slack.com/blablah go run *.go --render > alertmanager.yml
	@cat -n alertmanager.yml
	@echo "Result:"
	@go run *.go --lint
