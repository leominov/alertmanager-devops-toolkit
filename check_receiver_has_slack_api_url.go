package main

import "fmt"

func init() {
	RegisterCheck("receiver_has_slack_api_url", CheckReceiverHasSlackApiURL)
}

func CheckReceiverHasSlackApiURL(a *AlertmanagerConfig) []error {
	var errs []error
	for _, receiver := range a.Receivers {
		for _, slackConfig := range receiver.SlackConfigs {
			if len(slackConfig.ApiURL) != 0 {
				errs = append(errs, fmt.Errorf("Found api_url in %s receiver", receiver.Name))
			}
		}
	}
	return errs
}
