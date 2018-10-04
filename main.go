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
	RenderTemplate   = flag.Bool("render", false, "Render template")
	RootTemplateFile = flag.String("render-template", ".alertmanager.tmpl.yml", "Template file to render")
	LintTemplate     = flag.Bool("lint", false, "Lint config")
	RootConfigFile   = flag.String("lint-config", "alertmanager.yml", "Configuration file to lint")
	ShowVersion      = flag.Bool("version", false, "Prints version and exit")

	Version = "1.0.0"
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
	res, err := GenerateTemplate(file, vars)
	if err != nil {
		return "", err
	}
	return res, nil
}

func lintConfig(file string) []error {
	var errs []error

	b, err := ioutil.ReadFile(file)
	if err != nil {
		return []error{err}
	}

	config := &AlertmanagerConfig{}
	if err := yaml.Unmarshal(b, &config); err != nil {
		return []error{err}
	}

	errs = append(errs, config.CheckRouteReceiver()...)
	errs = append(errs, config.CheckReceivers()...)
	errs = append(errs, config.CheckEmptyReceivers()...)
	errs = append(errs, config.CheckSlackChannels()...)
	errs = append(errs, config.CheckSlackApiURL()...)
	errs = append(errs, config.CheckWebhookURLs()...)
	errs = append(errs, config.CheckEmailTo()...)
	errs = append(errs, config.CheckSlackHttpConfigProxyURL()...)
	errs = append(errs, config.CheckWebhookHttpConfigProxyURL()...)
	errs = append(errs, config.CheckDefaultReceiver()...)

	return errs
}

func printsErrorArray(errs []error) {
	for _, err := range errs {
		fmt.Println(err)
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
			fmt.Println(err)
			return 1
		}
		fmt.Println(res)
	}
	if *LintTemplate {
		errs := lintConfig(*RootConfigFile)
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
