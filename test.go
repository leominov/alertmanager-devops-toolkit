package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type RouteTest struct {
	Receivers []string          `yaml:"receivers"`
	Labels    map[string]string `yaml:"labels"`
}

func (r *RouteTest) String() string {
	return fmt.Sprintf("%v == %s", r.Labels, strings.Join(r.Receivers, ","))
}

func (r *RouteTest) Test(config string) error {
	args := []string{
		"config",
		"routes",
		"test",
		fmt.Sprintf("--config.file=%s", config),
		fmt.Sprintf("--verify.receivers=%s", strings.Join(r.Receivers, ",")),
	}
	for k, v := range r.Labels {
		args = append(args, fmt.Sprintf("%s=%s", k, v))
	}
	b, err := exec.Command("amtool", args...).CombinedOutput()
	out := strings.TrimSpace(string(b))
	if err != nil {
		return fmt.Errorf("%s: %s", r, out)
	}
	return nil
}

func checkForExists(items []string) error {
	for _, item := range items {
		if _, err := os.Stat(item); os.IsNotExist(err) {
			return err
		}
	}
	return nil
}

func isRouteTestsFile(info os.FileInfo) bool {
	if info.IsDir() {
		return false
	}
	if !strings.HasSuffix(info.Name(), "yaml") && !strings.HasSuffix(info.Name(), "yml") {
		return false
	}
	return true
}

func loadRouteTests(testDir string) ([]*RouteTest, error) {
	tests := []*RouteTest{}
	err := filepath.Walk(testDir, func(path string, info os.FileInfo, err error) error {
		if !isRouteTestsFile(info) {
			return nil
		}
		b, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		fileTests := []*RouteTest{}
		err = yaml.Unmarshal(b, &fileTests)
		if err != nil {
			return err
		}
		tests = append(tests, fileTests...)
		return nil
	})
	return tests, err
}

func RoutesTest(config string, testDir string) []error {
	if err := checkForExists([]string{config, testDir}); err != nil {
		return []error{err}
	}
	tests, err := loadRouteTests(testDir)
	if err != nil {
		return []error{err}
	}
	errs := []error{}
	for _, test := range tests {
		err := test.Test(config)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}
