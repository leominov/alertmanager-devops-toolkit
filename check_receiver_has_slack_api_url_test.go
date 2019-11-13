package main

import "testing"

func TestCheckReceiverHasSlackApiURL(t *testing.T) {
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
	errs := CheckReceiverHasSlackApiURL(a)
	if len(errs) != 0 {
		t.Error("CheckReceiverHasSlackApiURL() != 0")
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
	errs = CheckReceiverHasSlackApiURL(a)
	if len(errs) != 1 {
		t.Error("CheckReceiverHasSlackApiURL() != 1")
	}
}
