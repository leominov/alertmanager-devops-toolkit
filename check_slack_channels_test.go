package main

import "testing"

func TestCheckSlackChannels(t *testing.T) {
	a := &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
			},
		},
	}
	errs := CheckSlackChannels(a)
	if len(errs) != 0 {
		t.Error("CheckSlackChannels() != 0")
	}
	a = &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
				SlackConfigs: []*SlackConfig{
					{
						Channel: "",
					},
				},
			},
		},
	}
	errs = CheckSlackChannels(a)
	if len(errs) != 0 {
		t.Error("CheckSlackChannels() != 0")
	}
	a = &AlertmanagerConfig{
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
	errs = CheckSlackChannels(a)
	if len(errs) != 0 {
		t.Error("CheckSlackChannels() != 0")
	}
	a = &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
				SlackConfigs: []*SlackConfig{
					{
						Channel: "+l.aminov",
					},
				},
			},
		},
	}
	errs = CheckSlackChannels(a)
	if len(errs) != 1 {
		t.Error("CheckSlackChannels() != 1")
	}
}
