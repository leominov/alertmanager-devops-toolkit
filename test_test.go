package main

import (
	"os"
	"testing"
)

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
	for id, test := range tests {
		errs := RoutesTest(test.config, test.dir, []string{"*.yaml", "*.yml"})
		if len(errs) != test.errors {
			t.Errorf("%d. len(RoutesTest()) == %d, want: %d", id, len(errs), test.errors)
		}
	}
}

func TestIsRouteTestsFile(t *testing.T) {
	tests := []struct {
		file      string
		testFiles []string
		want      bool
	}{
		{
			file: ".alertmanager.tmpl.yml",
			testFiles: []string{
				"*.yml",
			},
			want: true,
		},
		{
			file: ".alertmanager.tmpl.yml",
			testFiles: []string{
				"*.tmpl.yml",
			},
			want: true,
		},
		{
			file: ".alertmanager.tmpl.yml",
			testFiles: []string{
				"*.foobar",
				"*.tmpl.yml",
			},
			want: true,
		},
		{
			file: ".alertmanager.tmpl.yml",
			testFiles: []string{
				"*.foobar",
			},
			want: false,
		},
		{
			file:      "tests",
			testFiles: []string{},
			want:      false,
		},
		{
			file:      ".alertmanager.tmpl.yml",
			testFiles: nil,
			want:      false,
		},
	}
	for _, test := range tests {
		info, err := os.Stat(test.file)
		if err != nil {
			t.Fatal(err)
		}
		result := IsRouteTestsFile(info, test.testFiles)
		if result != test.want {
			t.Errorf("IsRouteTestsFile() = %v, want = %v", result, test.want)
		}
	}
}
