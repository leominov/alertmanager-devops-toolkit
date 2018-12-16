export SLACK_API_URL=http://slack.com/blablah
export SMTP_HOST=localhost
export SMTP_FROM=no-reply@localhost.com
export SMTP_AUTH_USERNAME=user
export SMTP_AUTH_PASSWORD=pass

.PHONY: release lint
release:
	@./.release.sh

lint:
	@go build -o .build/alertmanager-devops-toolkit
	@./.build/alertmanager-devops-toolkit --render --safe=false | cat -n
	@./.build/alertmanager-devops-toolkit --render > alertmanager.yml
	@echo "Result:"
	@yamllint -s alertmanager.yml
	@./.build/alertmanager-devops-toolkit --lint
