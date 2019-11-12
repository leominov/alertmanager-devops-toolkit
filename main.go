package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

const (
	BlockTemplate = "#\n# File: %s\n%s\n"
)

var (
	RenderTemplate     = flag.Bool("render", false, "Render template")
	SafeRenderTemplate = flag.Bool("safe", true, "Included all specified secrets")
	RootTemplateFile   = flag.String("render-template", ".alertmanager.tmpl.yml", "Template file to render")
	LintTemplate       = flag.Bool("lint", false, "Lint config")
	LintConfigFile     = flag.String("lint-config", "alertmanager.yml", "Configuration file to lint")
	TestRoutes         = flag.Bool("test", false, "Test config")
	TestConfigFile     = flag.String("test-config", "alertmanager.yml", "Configuration file to test")
	TestDir            = flag.String("test-dir", "tests", "Directory with config tests")
	ShowVersion        = flag.Bool("version", false, "Prints version and exit")

	Version = "1.3.0"
)

func templateVars() (map[string]interface{}, error) {
	vars, err := ValuesFromDirectory("./")
	if err != nil {
		return nil, err
	}
	envs := ValuesFromEnviron()
	return map[string]interface{}{
		"Values": vars,
		"Env":    envs,
	}, nil
}

func renderTemplate(file string) (string, error) {
	vars, err := templateVars()
	if err != nil {
		return "", err
	}
	res, err := Render(file, vars)
	if err != nil {
		return "", err
	}
	return res, nil
}

func loadConfig(file string) (*AlertmanagerConfig, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	config := &AlertmanagerConfig{}
	if err := yaml.Unmarshal(b, &config); err != nil {
		return nil, err
	}
	return config, nil
}

func printsErrorArray(errs []error) {
	for _, err := range errs {
		fmt.Fprintf(os.Stderr, "[x] %v\n", err)
	}
}

func realMain() int {
	flag.Parse()
	if *ShowVersion {
		fmt.Println(Version)
		return 0
	}
	if *RenderTemplate {
		res, err := renderTemplate(*RootTemplateFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		fmt.Println(res)
	}
	if *LintTemplate {
		config, err := loadConfig(*LintConfigFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 2
		}
		errs := config.Lint()
		if len(errs) != 0 {
			printsErrorArray(errs)
			return 2
		}
		fmt.Println("Looks good to me")
	}
	if *TestRoutes {
		errs := RoutesTest(*TestConfigFile, *TestDir)
		if len(errs) != 0 {
			printsErrorArray(errs)
			return 2
		}
		fmt.Println("Looks good to me")
	}
	return 0
}

func main() {
	os.Exit(realMain())
}
