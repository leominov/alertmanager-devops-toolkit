package main

import "testing"

func TestCheckSlackApiURL(t *testing.T) {
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
	errs := CheckSlackApiURL(a)
	if len(errs) != 0 {
		t.Error("CheckSlackApiURL() != 0")
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
	errs = CheckSlackApiURL(a)
	if len(errs) != 1 {
		t.Error("CheckSlackApiURL() != 1")
	}
}
