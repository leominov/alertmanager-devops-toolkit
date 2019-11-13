package main

import "testing"

func TestCheckReceiverWebhookURLs(t *testing.T) {
	a := &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
				WebhookConfigs: []*WebhookConfig{
					{
						URL: "https://google.com",
					},
				},
			},
		},
	}
	errs := CheckReceiverWebhookURLs(a)
	if len(errs) != 0 {
		t.Error("CheckReceiverWebhookURLs() != 0")
	}
	a = &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
				WebhookConfigs: []*WebhookConfig{
					{
						URL: "http://[fe80::1%en0]/",
					},
				},
			},
		},
	}
	errs = CheckReceiverWebhookURLs(a)
	if len(errs) != 1 {
		t.Error("CheckReceiverWebhookURLs() != 1")
	}
}
