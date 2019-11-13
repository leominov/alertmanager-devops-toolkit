package main

import "fmt"

func init() {
	RegisterCheck("empty_receivers", CheckEmptyReceivers)
}

func CheckEmptyReceivers(a *AlertmanagerConfig) []error {
	var errs []error
	for _, receiver := range a.Receivers {
		if receiver.EmailConfigs == nil && receiver.SlackConfigs == nil && receiver.WebhookConfigs == nil {
			errs = append(errs, fmt.Errorf("Empty config in %s receiver", receiver.Name))
		}
	}
	return errs
}
