package main

import "testing"

func TestCheckWebhookHttpConfigProxyURL(t *testing.T) {
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
	errs := CheckWebhookHttpConfigProxyURL(a)
	if len(errs) != 0 {
		t.Error("CheckWebhookHttpConfigProxyURL() != 0")
	}
	a = &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
				WebhookConfigs: []*WebhookConfig{
					{
						URL: "https://google.com",
						HttpConfig: &HttpConfig{
							ProxyURL: "http://[fe80::1%en0]/",
						},
					},
				},
			},
		},
	}
	errs = CheckWebhookHttpConfigProxyURL(a)
	if len(errs) != 1 {
		t.Error("CheckWebhookHttpConfigProxyURL() != 1")
	}
}
