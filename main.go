package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/mail"
	"net/url"
	"os"
	"path"
	"path/filepath"
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

	funcMap = template.FuncMap{
		"indent":  indentTmpl,
		"include": includeTmpl,
	}

	Version = "1.0.0"
)

type AlertmanagerConfig struct {
	Receivers []*Receiver `yaml:"receivers"`
	RouteRoot *RouteRoot  `yaml:"route"`
}

type Receiver struct {
	Name           string           `yaml:"name"`
	SlackConfigs   []*SlackConfig   `yaml:"slack_configs"`
	EmailConfigs   []*EmailConfig   `yaml:"email_configs"`
	WebhookConfigs []*WebhookConfig `yaml:"webhook_configs"`
}

type WebhookConfig struct {
	URL        string      `yaml:"url"`
	HttpConfig *HttpConfig `yaml:"http_config"`
}

type EmailConfig struct {
	To string `yaml:"to"`
}

type SlackConfig struct {
	ApiURL     string      `yaml:"api_url"`
	Channel    string      `yaml:""channel`
	HttpConfig *HttpConfig `yaml:"http_config"`
}

type HttpConfig struct {
	ProxyURL string `yaml:"proxy_url"`
}

type RouteRoot struct {
	Receiver string   `yaml:"receiver"`
	Routes   []*Route `yaml:"routes"`
}

type Route struct {
	Receiver string `yaml:"receiver"`
}

func normalizeTextBlock(text string) string {
	text = strings.Replace(text, "---\n", "", -1)
	text = strings.TrimRight(text, "\n")
	return text
}

func indentTmpl(indent int, text string) string {
	pad := strings.Repeat(" ", indent)
	return pad + strings.Replace(text, "\n", "\n"+pad, -1)
}

func includeTmpl(name string, data map[string]interface{}) (string, error) {
	var result string
	files, err := filepath.Glob(name)
	if err != nil {
		return "", err
	}
	// Can't use template.ParseFiles file names can be same
	for _, file := range files {
		vars, err := LoadValuesFromDirectory(path.Dir(file))
		if err == nil {
			data["Group"] = vars
		}
		buf := bytes.NewBufferString("")
		tmpl, err := template.ParseFiles([]string{file}...)
		if err != nil {
			return "", err
		}
		if err := tmpl.Option(TemplateOption).Execute(buf, data); err != nil {
			return "", err
		}
		out := fmt.Sprintf(BlockTemplate, file, normalizeTextBlock(buf.String()))
		result += out
	}
	result = normalizeTextBlock(result)
	return result, nil
}

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

func LoadValuesFromDirectory(dir string) (map[string]interface{}, error) {
	var vars map[string]interface{}
	filePath := path.Join(dir, "values.yml")
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return vars, err
	}
	if err := yaml.Unmarshal(b, &vars); err != nil {
		return vars, err
	}
	return vars, nil
}

func EnvironToMap(envs []string) map[string]string {
	result := make(map[string]string)
	for _, env := range envs {
		data := strings.Split(env, "=")
		if len(data) < 2 {
			continue
		}
		key := data[0]
		val := data[1]
		result[key] = val
	}
	return result
}

func renderTemplate(file string) (string, error) {
	vars, err := LoadValuesFromDirectory("./")
	if err != nil {
		return "", err
	}
	envs := EnvironToMap(os.Environ())
	res, err := GenerateTemplate(file, map[string]interface{}{
		"Values": vars,
		"Env":    envs,
	})
	if err != nil {
		return "", err
	}
	return res, nil
}

func (a *AlertmanagerConfig) CheckReceivers() []error {
	var errs []error
	routeReceivers := make(map[string]bool)
	// Add default receiver
	routeReceivers[a.RouteRoot.Receiver] = true
	for _, route := range a.RouteRoot.Routes {
		routeReceivers[route.Receiver] = true
	}
	// All receivers must be used
	for _, receiver := range a.Receivers {
		if _, ok := routeReceivers[receiver.Name]; !ok {
			errs = append(errs, fmt.Errorf("Receiver %s does't found in any route", receiver.Name))
		}
	}
	return errs
}

func (a *AlertmanagerConfig) CheckSlackChannels() []error {
	var errs []error
	for _, receiver := range a.Receivers {
		for _, slackConfig := range receiver.SlackConfigs {
			ch := slackConfig.Channel
			if len(ch) == 0 {
				continue
			}
			if !strings.HasPrefix(ch, "#") && !strings.HasPrefix(ch, "@") {
				errs = append(errs, fmt.Errorf("Incorrent Slack channel in %s receiver: %s", receiver.Name, ch))
			}
		}
	}
	return errs
}

func (a *AlertmanagerConfig) CheckEmptyReceivers() []error {
	var errs []error
	for _, receiver := range a.Receivers {
		if receiver.EmailConfigs == nil && receiver.SlackConfigs == nil && receiver.WebhookConfigs == nil {
			errs = append(errs, fmt.Errorf("Empty config in %s receiver", receiver.Name))
		}
	}
	return errs
}

func (a *AlertmanagerConfig) CheckWebhookURLs() []error {
	var errs []error
	for _, receiver := range a.Receivers {
		for _, webhookConfig := range receiver.WebhookConfigs {
			_, err := url.Parse(webhookConfig.URL)
			if err != nil {
				errs = append(errs, fmt.Errorf("Receiver %s error with %s: %v", receiver.Name, webhookConfig.URL, err))
			}
		}
	}
	return errs
}

func (a *AlertmanagerConfig) CheckEmailTo() []error {
	var errs []error
	for _, receiver := range a.Receivers {
		for _, emailConfig := range receiver.EmailConfigs {
			recipients := strings.Split(emailConfig.To, ",")
			if len(recipients) == 0 {
				continue
			}
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

func (a *AlertmanagerConfig) CheckRouteReceiver() []error {
	var errs []error
	for id, route := range a.RouteRoot.Routes {
		if len(route.Receiver) == 0 {
			errs = append(errs, fmt.Errorf("Route #%d has empty receiver", id))
		}
	}
	return errs
}

func (a *AlertmanagerConfig) CheckSlackApiURL() []error {
	var errs []error
	for _, receiver := range a.Receivers {
		for _, slackConfig := range receiver.SlackConfigs {
			if len(slackConfig.ApiURL) != 0 {
				errs = append(errs, fmt.Errorf("Found api_url in %s receiver", receiver.Name))
			}
		}
	}
	return errs
}

func (a *AlertmanagerConfig) CheckSlackHttpConfigProxyURL() []error {
	var errs []error
	for _, receiver := range a.Receivers {
		for _, slackConfig := range receiver.SlackConfigs {
			if slackConfig.HttpConfig == nil {
				continue
			}
			_, err := url.Parse(slackConfig.HttpConfig.ProxyURL)
			if err != nil {
				errs = append(errs, fmt.Errorf("Receiver %s error with %s: %v", receiver.Name, slackConfig.HttpConfig.ProxyURL, err))
			}
		}
	}
	return errs
}

func (a *AlertmanagerConfig) CheckWebhookHttpConfigProxyURL() []error {
	var errs []error
	for _, receiver := range a.Receivers {
		for _, webhookConfig := range receiver.WebhookConfigs {
			if webhookConfig.HttpConfig == nil {
				continue
			}
			_, err := url.Parse(webhookConfig.HttpConfig.ProxyURL)
			if err != nil {
				errs = append(errs, fmt.Errorf("Receiver %s error with %s: %v", receiver.Name, webhookConfig.HttpConfig.ProxyURL, err))
			}
		}
	}
	return errs
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
