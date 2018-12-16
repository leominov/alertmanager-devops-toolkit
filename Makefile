export SLACK_API_URL=http://slack.com/blablah
export SMTP_HOST=localhost
export SMTP_FROM=no-reply@localhost.com
export SMTP_AUTH_USERNAME=user
export SMTP_AUTH_PASSWORD=pass

.PHONY: release lint
release:
	@./.release.sh

lint:
	@go run *.go --render --safe=false | cat -n
	@go run *.go --render > alertmanager.yml
	@echo "Result:"
	@yamllint -s alertmanager.yml
	@go run *.go --lint
