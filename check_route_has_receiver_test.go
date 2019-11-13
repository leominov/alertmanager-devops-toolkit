package main

import "testing"

func TestCheckRouteHasReceiver(t *testing.T) {
	a := &AlertmanagerConfig{
		RouteRoot: &RouteRoot{
			Routes: []*Route{
				{
					Receiver: "foobar",
				},
			},
		},
		Receivers: []*Receiver{
			{
				Name: "foobar",
			},
		},
	}
	errs := CheckRouteHasReceiver(a)
	if len(errs) != 0 {
		t.Error("CheckRouteHasReceiver() != 0")
	}
	a = &AlertmanagerConfig{
		RouteRoot: &RouteRoot{
			Routes: []*Route{
				{
					Receiver: "",
				},
			},
		},
		Receivers: []*Receiver{
			{
				Name: "foobar",
			},
		},
	}
	errs = CheckRouteHasReceiver(a)
	if len(errs) != 1 {
		t.Error("CheckRouteHasReceiver() != 1")
	}
}
