package main

import "fmt"

func init() {
	RegisterCheck("receiver_unique_email_to", CheckReceiverUniqueEmailTo)
}

func CheckReceiverUniqueEmailTo(a *AlertmanagerConfig) []error {
	var errs []error
	for _, receiver := range a.Receivers {
		emails := make(map[string]bool)
		if len(receiver.EmailConfigs) == 0 {
			continue
		}
		for _, emailConfig := range receiver.EmailConfigs {
			_, ok := emails[emailConfig.To]
			if ok {
				errs = append(errs, fmt.Errorf("Non-unique Email %s in %s receiver", emailConfig.To, receiver.Name))
			}
			emails[emailConfig.To] = true
		}
	}
	return errs
}
