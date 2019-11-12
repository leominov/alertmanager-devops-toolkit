# Alertmanager DevOps Toolkit

## Tool #1. Linter

### Rules

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

Render given template, `include` and `indent` are supported, useful for categorization of receivers and routes (see [GitLab Code Owners](https://docs.gitlab.com/ee/user/project/code_owners.html) and [Atlassian Stash Approvers](https://github.com/leominov/atlas-hook#prwizard)). For example, you have an structure:

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
    Included all specified secrets (default true)
```

## Tool #3. Tests

Wrapper for `amtool config routes test` command, arguments was stored in YAML-files.

```
tests/
    groupA.yaml
```

`groupA.yaml` content:

```
---
- receivers:
    - group1-dev
  labels:
    env: dev
```

### Usage

```
-test
    Test config
-test-config string
    Configuration file to test (default "alertmanager.yml")
-test-dir string
    Directory with config tests (default "tests")
```
