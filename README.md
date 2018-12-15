# Alertmanager DevOps Toolkit

## Tool #1. Linter

* All routes must contains receiver;
* All receivers must be used;
* Receivers must contains `slack_configs`, `email_configs` or `webhook_configs`;
* Slack `api_url` only in global settings;
* Slack channels starts with `#` or `@`;
* Only valid webhook URLs;
* Only valid Email recipients.

And some other.

### Usage

```
-lint
    Lint config
-lint-config string
    Configuration file to lint (default "alertmanager.yml")
```

## Tool #2. Render

Render given template, `include` and `indent` are supported, useful for categorization of receivers and routes. For example, you have an structure:

```
conf/
    teamA/
        routes.yml
        receivers.yml
    teamB/
        routes.yml
        receivers.yml
    teamC/
        routes.yml
        receivers.yml
```

And template of configuration file:

```
---
global:
  slack_api_url: {{ .Env.SLACK_API_URL | secret }}
receivers:
{{ include "conf/*/receivers.yml" . | indent 2 }}
route:
  receiver: receiver
  routes:
{{ include "conf/*/routes.yml" . | indent 4 }}
```

After rendering you got:

```
---
global:
  slack_api_url: http://slack.com/blablah
receivers:
  # conf/teamA/receivers.yml content
  ...
  # conf/teamB/receivers.yml content
  ...
  # conf/teamC/receivers.yml content
  ...
route:
  receiver: receiver
  routes:
    # conf/teamA/routes.yml content
    ...
    # conf/teamB/routes.yml content
    ...
    # conf/teamC/routes.yml content
    ...
```

### Usage

```
-render
    Render template
-render-template string
    Template file to render (default ".alertmanager.tmpl.yml")
  -safe
    Render template with all included secrets (default true)
```
