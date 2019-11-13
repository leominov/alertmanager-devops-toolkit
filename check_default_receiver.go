package main

import "errors"

func CheckDefaultReceiver(a *AlertmanagerConfig) []error {
	for _, receiver := range a.Receivers {
		if receiver.Name == a.RouteRoot.Receiver {
			return nil
		}
	}
	return []error{
		errors.New("Default receiver doesn't found in list"),
	}
}
