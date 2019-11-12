package main

import "testing"

func TestRoutesTest(t *testing.T) {
	tests := []struct {
		config, dir string
		errors      int
	}{
		{
			config: "conf/alertmanager.yml",
			dir:    "tests",
			errors: 2,
		},
		{
			config: "conf/not-found.yml",
			dir:    "tests",
			errors: 1,
		},
		{
			config: "conf/alertmanager.yml",
			dir:    "not-found",
			errors: 1,
		},
		{
			config: "conf/alertmanager.yml",
			dir:    "tests_invalid",
			errors: 1,
		},
	}
	for _, test := range tests {
		errs := RoutesTest(test.config, test.dir)
		if len(errs) != test.errors {
			t.Errorf("len(RoutesTest()) == %d, want: %d", len(errs), test.errors)
		}
	}
}
