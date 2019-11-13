package main

import "testing"

func TestCheckReceiverEmailTo(t *testing.T) {
	a := &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
				EmailConfigs: []*EmailConfig{
					{
						To: "",
					},
				},
			},
		},
	}
	errs := CheckReceiverEmailTo(a)
	if len(errs) != 0 {
		t.Error("CheckReceiverEmailTo() != 0")
	}
	a = &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
				EmailConfigs: []*EmailConfig{
					{
						To: "foobar@gmail.com",
					},
				},
			},
		},
	}
	errs = CheckReceiverEmailTo(a)
	if len(errs) != 0 {
		t.Error("CheckReceiverEmailTo() != 0")
	}
	a = &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
				EmailConfigs: []*EmailConfig{
					{
						To: "Joe Doe",
					},
				},
			},
		},
	}
	errs = CheckReceiverEmailTo(a)
	if len(errs) != 1 {
		t.Error("CheckReceiverEmailTo() != 1")
	}
}
