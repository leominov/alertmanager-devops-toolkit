package main

import "testing"

func TestCheckSlackHttpConfigProxyURL(t *testing.T) {
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
	errs := CheckSlackHttpConfigProxyURL(a)
	if len(errs) != 0 {
		t.Error("CheckSlackHttpConfigProxyURL() != 0")
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
	errs = CheckSlackHttpConfigProxyURL(a)
	if len(errs) != 1 {
		t.Error("CheckSlackHttpConfigProxyURL() != 1")
	}
}
