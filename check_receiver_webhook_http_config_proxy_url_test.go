package main

import "testing"

func TestCheckReceiverWebhookHttpConfigProxyURL(t *testing.T) {
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
	errs := CheckReceiverWebhookHttpConfigProxyURL(a, defaultCheckOptions)
	if len(errs) != 0 {
		t.Error("CheckReceiverWebhookHttpConfigProxyURL() != 0")
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
	errs = CheckReceiverWebhookHttpConfigProxyURL(a, defaultCheckOptions)
	if len(errs) != 1 {
		t.Error("CheckReceiverWebhookHttpConfigProxyURL() != 1")
	}
}
