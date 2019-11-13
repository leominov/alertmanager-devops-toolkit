package main

import "fmt"

func init() {
	RegisterCheck("route_receiver_is_defined", CheckRouteReceiverIsDefined)
}

func CheckRouteReceiverIsDefined(a *AlertmanagerConfig) []error {
	var errs []error
	routeReceivers := make(map[string]bool)
	// Add default receiver
	routeReceivers[a.RouteRoot.Receiver] = true
	for _, route := range a.RouteRoot.Routes {
		routeReceivers[route.Receiver] = true
	}
	// All receivers must be used
	for _, receiver := range a.Receivers {
		if _, ok := routeReceivers[receiver.Name]; !ok {
			errs = append(errs, fmt.Errorf("Receiver %s wasn't found in any route", receiver.Name))
		}
	}
	return errs
}
