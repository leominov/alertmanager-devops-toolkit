package main

import "testing"

func TestCheckDefaultReceiver(t *testing.T) {
	a := &AlertmanagerConfig{
		RouteRoot: &RouteRoot{
			Receiver: "foobar",
		},
		Receivers: []*Receiver{
			{
				Name: "foobar",
			},
		},
	}
	errs := CheckDefaultReceiver(a)
	if len(errs) != 0 {
		t.Error("CheckDefaultReceiver() != 0")
	}
	a = &AlertmanagerConfig{
		RouteRoot: &RouteRoot{
			Receiver: "barfoo",
		},
		Receivers: []*Receiver{
			{
				Name: "foobar",
			},
		},
	}
	errs = CheckDefaultReceiver(a)
	if len(errs) != 1 {
		t.Error("CheckDefaultReceiver() != 1")
	}
}
