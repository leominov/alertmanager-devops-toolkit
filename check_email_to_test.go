package main

import "testing"

func TestCheckEmailTo(t *testing.T) {
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
	errs := CheckEmailTo(a)
	if len(errs) != 0 {
		t.Error("CheckEmailTo() != 0")
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
	errs = CheckEmailTo(a)
	if len(errs) != 0 {
		t.Error("CheckEmailTo() != 0")
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
	errs = CheckEmailTo(a)
	if len(errs) != 1 {
		t.Error("CheckEmailTo() != 1")
	}
}
