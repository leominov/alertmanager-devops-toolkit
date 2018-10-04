# Alertmanager DevOps Toolkit

## Tool #1. Linter

* All routes must contains receiver;
* All receivers must be used;
* Receivers must contains `slack_configs`, `email_configs` or `webhook_configs`;
* Slack `api_url` only in global settings;
* Slack channels starts with `#` or `@`;
* Only valid webhook URLs;
* Only valid Email recipients.

### Usage

```
-lint
    Lint config
-lint-config string
    Configuration file to lint (default "alertmanager.yml")
```

## Tool #2. Render

Render given template, `include` and `indent` are supported, useful for categorization of receivers and routes.

### Usage

```
-render
    Render template
-render-template string
    Template file to render (default ".alertmanager.tmpl.yml")
```
