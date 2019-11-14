package main

import "fmt"

func init() {
	RegisterCheck("route_has_receiver", CheckRouteHasReceiver)
}

func CheckRouteHasReceiver(a *AlertmanagerConfig, opt *CheckOptions) []error {
	var errs []error
	for id, route := range a.RouteRoot.Routes {
		if len(route.Receiver) == 0 {
			errs = append(errs, fmt.Errorf("Route #%d has empty receiver", id))
		}
	}
	return errs
}
