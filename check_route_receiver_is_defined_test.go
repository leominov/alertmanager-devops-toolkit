package main

import "testing"

func TestCheckRouteReceiverIsDefined(t *testing.T) {
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
	errs := CheckRouteReceiverIsDefined(a, defaultCheckOptions)
	if len(errs) != 0 {
		t.Error("CheckRouteReceiverIsDefined() != 0")
	}
	a = &AlertmanagerConfig{
		RouteRoot: &RouteRoot{
			Routes: []*Route{
				{
					Receiver: "foobar",
				},
			},
		},
		Receivers: []*Receiver{
			{
				Name: "barfoo",
			},
		},
	}
	errs = CheckRouteReceiverIsDefined(a, defaultCheckOptions)
	if len(errs) != 1 {
		t.Error("CheckRouteReceiverIsDefined() != 1")
	}
}
