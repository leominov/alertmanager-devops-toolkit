package main

type AlertmanagerConfig struct {
	Receivers []*Receiver `yaml:"receivers"`
	RouteRoot *RouteRoot  `yaml:"route"`
}

type Receiver struct {
	Name           string           `yaml:"name"`
	SlackConfigs   []*SlackConfig   `yaml:"slack_configs"`
	EmailConfigs   []*EmailConfig   `yaml:"email_configs"`
	WebhookConfigs []*WebhookConfig `yaml:"webhook_configs"`
}

type WebhookConfig struct {
	URL        string      `yaml:"url"`
	HttpConfig *HttpConfig `yaml:"http_config"`
}

type EmailConfig struct {
	To string `yaml:"to"`
}

type SlackConfig struct {
	ApiURL     string      `yaml:"api_url"`
	Channel    string      `yaml:""channel`
	HttpConfig *HttpConfig `yaml:"http_config"`
}

type HttpConfig struct {
	ProxyURL string `yaml:"proxy_url"`
}

type RouteRoot struct {
	Receiver string   `yaml:"receiver"`
	Routes   []*Route `yaml:"routes"`
}

type Route struct {
	Receiver string `yaml:"receiver"`
}

type StringSlice []string

func (s *StringSlice) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var stringType string
	if err := unmarshal(&stringType); err == nil {
		*s = []string{stringType}
		return nil
	}

	var sliceType []string
	if err := unmarshal(&sliceType); err != nil {
		return err
	}
	*s = sliceType
	return nil
}
