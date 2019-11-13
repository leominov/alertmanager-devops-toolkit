package main

import "testing"

func TestList(t *testing.T) {
	config := &AlertmanagerConfig{
		RouteRoot: &RouteRoot{},
	}
	Lint(config)
}
