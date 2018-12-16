export SLACK_API_URL=http://slack.com/blablah
export SMTP_HOST=localhost
export SMTP_FROM=no-reply@localhost.com
export SMTP_AUTH_USERNAME=user
export SMTP_AUTH_PASSWORD=pass

.PHONY: release lint
release:
	@./.release.sh

lint:
	@go build -o alertmanager-devops-toolkit
	@./alertmanager-devops-toolkit --render --safe=false | cat -n
	@./alertmanager-devops-toolkit --render > alertmanager.yml
	@echo "Result:"
	@yamllint -s alertmanager.yml
	@./alertmanager-devops-toolkit --lint
