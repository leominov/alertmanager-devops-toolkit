# Alertmanager DevOps Toolkit

## Linters

* All routes must contains receiver;
* All receivers must be used;
* Receivers must contains `slack_configs`, `email_configs` or `webhook_configs`;
* Slack `api_url` only in global settings;
* Slack channels starts with `#` or `@`;
* Only valid webhook URLs;
* Only valid Email recipients.

## Usage

```
-lint
    Lint config
-lint-config string
    Configuration file to lint (default "alertmanager.yml")
-render
    Render template
-render-template string
    Template file to render (default ".alertmanager.tmpl.yml")
-version
    Prints version and exit
```
