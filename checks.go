package main

import (
	"errors"
	"fmt"
	"net/mail"
	"net/url"
	"strings"
)

func (a *AlertmanagerConfig) CheckReceivers() []error {
	var errs []error
	routeReceivers := make(map[string]bool)
	// Add default receiver
	routeReceivers[a.RouteRoot.Receiver] = true
	for _, route := range a.RouteRoot.Routes {
		routeReceivers[route.Receiver] = true
	}
	// All receivers must be used
	for _, receiver := range a.Receivers {
		if _, ok := routeReceivers[receiver.Name]; !ok {
			errs = append(errs, fmt.Errorf("Receiver %s does't found in any route", receiver.Name))
		}
	}
	return errs
}

func (a *AlertmanagerConfig) CheckSlackChannels() []error {
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

func (a *AlertmanagerConfig) CheckEmptyReceivers() []error {
	var errs []error
	for _, receiver := range a.Receivers {
		if receiver.EmailConfigs == nil && receiver.SlackConfigs == nil && receiver.WebhookConfigs == nil {
			errs = append(errs, fmt.Errorf("Empty config in %s receiver", receiver.Name))
		}
	}
	return errs
}

func (a *AlertmanagerConfig) CheckWebhookURLs() []error {
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

func (a *AlertmanagerConfig) CheckEmailTo() []error {
	var errs []error
	for _, receiver := range a.Receivers {
		for _, emailConfig := range receiver.EmailConfigs {
			recipients := strings.Split(emailConfig.To, ",")
			if len(recipients) == 0 {
				continue
			}
			for _, recipient := range recipients {
				recipient = strings.TrimSpace(recipient)
				_, err := mail.ParseAddress(recipient)
				if err != nil {
					errs = append(errs, fmt.Errorf("Receiver %s error with %s: %v", receiver.Name, recipient, err))
				}
			}
		}
	}
	return errs
}

func (a *AlertmanagerConfig) CheckRouteReceiver() []error {
	var errs []error
	for id, route := range a.RouteRoot.Routes {
		if len(route.Receiver) == 0 {
			errs = append(errs, fmt.Errorf("Route #%d has empty receiver", id))
		}
	}
	return errs
}

func (a *AlertmanagerConfig) CheckSlackApiURL() []error {
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

func (a *AlertmanagerConfig) CheckSlackHttpConfigProxyURL() []error {
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

func (a *AlertmanagerConfig) CheckWebhookHttpConfigProxyURL() []error {
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

func (a *AlertmanagerConfig) CheckDefaultReceiver() []error {
	for _, receiver := range a.Receivers {
		if receiver.Name == a.RouteRoot.Receiver {
			return nil
		}
	}
	return []error{
		errors.New("Default receiver doesn't found in list"),
	}
}
