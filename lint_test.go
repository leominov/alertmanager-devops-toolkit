package main

import "testing"

func TestList(t *testing.T) {
	aconfig := &AlertmanagerConfig{
		RouteRoot: &RouteRoot{},
	}
	Lint(aconfig, DefaultConfig())
	Lint(aconfig, &Config{
		Checks: map[string]*CheckOptions{
			"receiver_webhook_urls": &CheckOptions{
				Active: false,
			},
		},
	})
}
