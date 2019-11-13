package main

import "fmt"

func init() {
	RegisterCheck("receiver_unique_slack_channel", CheckReceiverUniqueSlackChannel)
}

func CheckReceiverUniqueSlackChannel(a *AlertmanagerConfig) []error {
	var errs []error
	for _, receiver := range a.Receivers {
		channels := make(map[string]bool)
		if len(receiver.SlackConfigs) == 0 {
			continue
		}
		for _, slackConfig := range receiver.SlackConfigs {
			_, ok := channels[slackConfig.Channel]
			if ok {
				errs = append(errs, fmt.Errorf("Non-unique Slack channel %s in %s receiver", slackConfig.Channel, receiver.Name))
			}
			channels[slackConfig.Channel] = true
		}
	}
	return errs
}
