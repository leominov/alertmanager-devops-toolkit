package main

import "fmt"

func CheckReceiverUniqueWebhookURL(a *AlertmanagerConfig) []error {
	var errs []error
	for _, receiver := range a.Receivers {
		links := make(map[string]bool)
		if len(receiver.WebhookConfigs) == 0 {
			continue
		}
		for _, webhookConfig := range receiver.WebhookConfigs {
			_, ok := links[webhookConfig.URL]
			if ok {
				errs = append(errs, fmt.Errorf("Non-unique webhook URL %s in %s receiver", webhookConfig.URL, receiver.Name))
			}
			links[webhookConfig.URL] = true
		}
	}
	return errs
}
