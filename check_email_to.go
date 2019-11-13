package main

import (
	"fmt"
	"net/mail"
	"strings"
)

func CheckEmailTo(a *AlertmanagerConfig) []error {
	var errs []error
	for _, receiver := range a.Receivers {
		for _, emailConfig := range receiver.EmailConfigs {
			if len(emailConfig.To) == 0 {
				continue
			}
			recipients := strings.Split(emailConfig.To, ",")
			for _, recipient := range recipients {
				recipient = strings.TrimSpace(recipient)
				_, err := mail.ParseAddress(recipient)
				if err != nil {
					errs = append(errs, fmt.Errorf("Receiver %s error with %s: %v", receiver.Name, recipient, err))
				}
			}
		}
	}
	return errs
}
