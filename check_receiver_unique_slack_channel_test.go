package main

import "testing"

func TestCheckReceiverUniqueSlackChannel(t *testing.T) {
	a := &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
				SlackConfigs: []*SlackConfig{
					{
						Channel: "@l.aminov",
					},
					{
						Channel: "@m.aminova",
					},
				},
			},
		},
	}
	errs := CheckReceiverUniqueSlackChannel(a, defaultCheckOptions)
	if len(errs) != 0 {
		t.Error("CheckReceiverUniqueSlackChannel() != 0")
	}
	a = &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
			},
		},
	}
	errs = CheckReceiverUniqueSlackChannel(a, defaultCheckOptions)
	if len(errs) != 0 {
		t.Error("CheckReceiverUniqueSlackChannel() != 0")
	}
	a = &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
				SlackConfigs: []*SlackConfig{
					{
						Channel: "@l.aminov",
					},
					{
						Channel: "@l.aminov",
					},
				},
			},
		},
	}
	errs = CheckReceiverUniqueSlackChannel(a, defaultCheckOptions)
	if len(errs) != 1 {
		t.Error("CheckReceiverUniqueSlackChannel() != 1")
	}
}
