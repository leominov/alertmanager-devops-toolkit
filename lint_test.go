package main

import "testing"

func TestList(t *testing.T) {
	a := &AlertmanagerConfig{
		RouteRoot: &RouteRoot{},
	}
	a.Lint()
}
