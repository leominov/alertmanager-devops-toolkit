package main

import (
	"fmt"
	"net/url"
)

func init() {
	RegisterCheck("receiver_webhook_urls", CheckReceiverWebhookURLs)
}

func CheckReceiverWebhookURLs(a *AlertmanagerConfig) []error {
	var errs []error
	for _, receiver := range a.Receivers {
		for _, webhookConfig := range receiver.WebhookConfigs {
			_, err := url.Parse(webhookConfig.URL)
			if err != nil {
				errs = append(errs, fmt.Errorf("Receiver %s error with %s: %v", receiver.Name, webhookConfig.URL, err))
			}
		}
	}
	return errs
}
