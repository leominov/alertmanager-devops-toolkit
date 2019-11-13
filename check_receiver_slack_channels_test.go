package main

import "testing"

func TestCheckReceiverSlackChannels(t *testing.T) {
	a := &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
			},
		},
	}
	errs := CheckReceiverSlackChannels(a)
	if len(errs) != 0 {
		t.Error("CheckReceiverSlackChannels() != 0")
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
	errs = CheckReceiverSlackChannels(a)
	if len(errs) != 0 {
		t.Error("CheckReceiverSlackChannels() != 0")
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
	errs = CheckReceiverSlackChannels(a)
	if len(errs) != 0 {
		t.Error("CheckReceiverSlackChannels() != 0")
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
	errs = CheckReceiverSlackChannels(a)
	if len(errs) != 1 {
		t.Error("CheckReceiverSlackChannels() != 1")
	}
}
