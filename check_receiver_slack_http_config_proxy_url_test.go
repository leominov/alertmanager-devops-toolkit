package main

import "testing"

func TestCheckReceiverSlackHttpConfigProxyURL(t *testing.T) {
	a := &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
				SlackConfigs: []*SlackConfig{
					{
						Channel: "@l.aminov",
					},
				},
			},
		},
	}
	errs := CheckReceiverSlackHttpConfigProxyURL(a, defaultCheckOptions)
	if len(errs) != 0 {
		t.Error("CheckReceiverSlackHttpConfigProxyURL() != 0")
	}
	a = &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
				SlackConfigs: []*SlackConfig{
					{
						Channel: "@l.aminov",
						HttpConfig: &HttpConfig{
							ProxyURL: "http://[fe80::1%en0]/",
						},
					},
				},
			},
		},
	}
	errs = CheckReceiverSlackHttpConfigProxyURL(a, defaultCheckOptions)
	if len(errs) != 1 {
		t.Error("CheckReceiverSlackHttpConfigProxyURL() != 1")
	}
}
