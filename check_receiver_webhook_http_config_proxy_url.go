package main

import (
	"fmt"
	"net/url"
)

func init() {
	RegisterCheck("receiver_webhook_http_config_proxy_url", CheckReceiverWebhookHttpConfigProxyURL)
}

func CheckReceiverWebhookHttpConfigProxyURL(a *AlertmanagerConfig, opt *CheckOptions) []error {
	var errs []error
	for _, receiver := range a.Receivers {
		for _, webhookConfig := range receiver.WebhookConfigs {
			if webhookConfig.HttpConfig == nil {
				continue
			}
			_, err := url.Parse(webhookConfig.HttpConfig.ProxyURL)
			if err != nil {
				errs = append(errs, fmt.Errorf("Receiver %s error with %s: %v", receiver.Name, webhookConfig.HttpConfig.ProxyURL, err))
			}
		}
	}
	return errs
}
