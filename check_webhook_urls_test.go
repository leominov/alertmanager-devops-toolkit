package main

import "testing"

func TestCheckWebhookURLs(t *testing.T) {
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
	errs := CheckWebhookURLs(a)
	if len(errs) != 0 {
		t.Error("CheckWebhookURLs() != 0")
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
	errs = CheckWebhookURLs(a)
	if len(errs) != 1 {
		t.Error("CheckWebhookURLs() != 1")
	}
}
