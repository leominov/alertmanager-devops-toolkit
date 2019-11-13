package main

import (
	"fmt"
	"net/url"
)

func init() {
	RegisterCheck("receiver_slack_http_config_proxy_url", CheckDefaultReceiver)
}

func CheckReceiverSlackHttpConfigProxyURL(a *AlertmanagerConfig) []error {
	var errs []error
	for _, receiver := range a.Receivers {
		for _, slackConfig := range receiver.SlackConfigs {
			if slackConfig.HttpConfig == nil {
				continue
			}
			_, err := url.Parse(slackConfig.HttpConfig.ProxyURL)
			if err != nil {
				errs = append(errs, fmt.Errorf("Receiver %s error with %s: %v", receiver.Name, slackConfig.HttpConfig.ProxyURL, err))
			}
		}
	}
	return errs
}
