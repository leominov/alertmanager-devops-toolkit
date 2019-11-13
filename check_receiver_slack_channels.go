package main

import (
	"fmt"
	"strings"
)

func CheckReceiverSlackChannels(a *AlertmanagerConfig) []error {
	var errs []error
	for _, receiver := range a.Receivers {
		for _, slackConfig := range receiver.SlackConfigs {
			ch := slackConfig.Channel
			if len(ch) == 0 {
				continue
			}
			if !strings.HasPrefix(ch, "#") && !strings.HasPrefix(ch, "@") {
				errs = append(errs, fmt.Errorf("Incorrent Slack channel in %s receiver: %s", receiver.Name, ch))
			}
		}
	}
	return errs
}
