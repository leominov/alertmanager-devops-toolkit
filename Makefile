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
	@amtool check-config alertmanager.yml
	@./.build/alertmanager-devops-toolkit --lint

.PHONY: test
test:
	@go test -cover

.PHONY: test-report
test-report:
	@go test -coverprofile=coverage.txt && go tool cover -html=coverage.txt
