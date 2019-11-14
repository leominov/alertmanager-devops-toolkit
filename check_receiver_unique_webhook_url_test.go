package main

import "testing"

func TestCheckReceiverUniqueWebhookURL(t *testing.T) {
	a := &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
				WebhookConfigs: []*WebhookConfig{
					{
						URL: "https://google.com",
					},
					{
						URL: "https://google.ru",
					},
				},
			},
		},
	}
	errs := CheckReceiverUniqueWebhookURL(a, defaultCheckOptions)
	if len(errs) != 0 {
		t.Error("CheckReceiverUniqueWebhookURL() != 0")
	}
	a = &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
			},
		},
	}
	errs = CheckReceiverUniqueWebhookURL(a, defaultCheckOptions)
	if len(errs) != 0 {
		t.Error("CheckReceiverUniqueWebhookURL() != 0")
	}
	a = &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
				WebhookConfigs: []*WebhookConfig{
					{
						URL: "https://google.com",
					},
					{
						URL: "https://google.com",
					},
				},
			},
		},
	}
	errs = CheckReceiverUniqueWebhookURL(a, defaultCheckOptions)
	if len(errs) != 1 {
		t.Error("CheckReceiverUniqueWebhookURL() != 1")
	}
}
