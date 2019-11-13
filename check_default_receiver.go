package main

import "errors"

func init() {
	RegisterCheck("default_receiver", CheckDefaultReceiver)
}

func CheckDefaultReceiver(a *AlertmanagerConfig) []error {
	for _, receiver := range a.Receivers {
		if receiver.Name == a.RouteRoot.Receiver {
			return nil
		}
	}
	return []error{
		errors.New("Default receiver wasn't found in list"),
	}
}
