package main

import "testing"

func TestCheckReceiverSlackApiURL(t *testing.T) {
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
	errs := CheckReceiverSlackApiURL(a)
	if len(errs) != 0 {
		t.Error("CheckReceiverSlackApiURL() != 0")
	}
	a = &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
				SlackConfigs: []*SlackConfig{
					{
						ApiURL: "https://google.com",
					},
				},
			},
		},
	}
	errs = CheckReceiverSlackApiURL(a)
	if len(errs) != 1 {
		t.Error("CheckReceiverSlackApiURL() != 1")
	}
}
