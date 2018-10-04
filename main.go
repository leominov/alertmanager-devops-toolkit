package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	yaml "gopkg.in/yaml.v2"
)

const (
	BlockTemplate  = "#\n# File: %s\n%s\n"
	TemplateOption = "missingkey=error"
)

var (
	RenderTemplate   = flag.Bool("render", false, "Render template")
	RootTemplateFile = flag.String("render-template", ".alertmanager.tmpl.yml", "Template file to render")
	LintTemplate     = flag.Bool("lint", false, "Lint config")
	RootConfigFile   = flag.String("lint-config", "alertmanager.yml", "Configuration file to lint")
	ShowVersion      = flag.Bool("version", false, "Prints version and exit")

	Version = "1.0.0"
)

func GenerateTemplate(filename string, data map[string]interface{}) (string, error) {
	buf := bytes.NewBufferString("")
	if tmpl, err := template.New(filename).Option(TemplateOption).Funcs(funcMap).ParseFiles([]string{filename}...); err != nil {
		return "", fmt.Errorf("failed to parse template for %s: %v", filename, err)
	} else if err := tmpl.Execute(buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template for %s: %v", filename, err)
	}
	result := strings.TrimRight(buf.String(), "\n")
	return result, nil
}

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
