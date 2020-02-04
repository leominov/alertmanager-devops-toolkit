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

.PHONY: test test-report
test:
	@go test -cover

test-report:
	@go test -coverprofile=coverage.txt && go tool cover -html=coverage.txt
