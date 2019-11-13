package main

import "testing"

func TestCheckEmptyReceivers(t *testing.T) {
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
	errs := CheckEmptyReceivers(a)
	if len(errs) != 0 {
		t.Error("CheckEmptyReceivers() != 0")
	}
	a = &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
			},
		},
	}
	errs = CheckEmptyReceivers(a)
	if len(errs) != 1 {
		t.Error("CheckEmptyReceivers() != 1")
	}
}
