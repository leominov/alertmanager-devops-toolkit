package main

import "testing"

func TestCheckReceiverUniqueEmailTo(t *testing.T) {
	a := &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
				EmailConfigs: []*EmailConfig{
					{
						To: "foobar@gmail.com",
					},
					{
						To: "barfoo@gmail.com",
					},
				},
			},
		},
	}
	errs := CheckReceiverUniqueEmailTo(a)
	if len(errs) != 0 {
		t.Error("CheckReceiverUniqueEmailTo() != 0")
	}
	a = &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
			},
		},
	}
	errs = CheckReceiverUniqueEmailTo(a)
	if len(errs) != 0 {
		t.Error("CheckReceiverUniqueEmailTo() != 0")
	}
	a = &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
				EmailConfigs: []*EmailConfig{
					{
						To: "foobar@gmail.com",
					},
					{
						To: "foobar@gmail.com",
					},
				},
			},
		},
	}
	errs = CheckReceiverUniqueEmailTo(a)
	if len(errs) != 1 {
		t.Error("CheckReceiverUniqueEmailTo() != 1")
	}
}
