package main

import (
	"bytes"
	"fmt"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	yaml "gopkg.in/yaml.v2"
)

var (
	funcMap = template.FuncMap{
		"indent":   Indent,
		"include":  Include,
		"toYaml":   ToYaml,
		"fromYaml": FromYaml,
	}
)

func ToYaml(v interface{}) string {
	data, err := yaml.Marshal(v)
	if err != nil {
		return ""
	}
	return string(data)
}

func FromYaml(str string) map[string]interface{} {
	m := map[string]interface{}{}
	if err := yaml.Unmarshal([]byte(str), &m); err != nil {
		m["Error"] = err.Error()
	}
	return m
}

func Indent(indent int, text string) string {
	pad := strings.Repeat(" ", indent)
	return pad + strings.Replace(text, "\n", "\n"+pad, -1)
}

func NormalizeTextBlock(text string) string {
	text = strings.Replace(text, "---\n", "", -1)
	text = strings.TrimRight(text, "\n")
	return text
}

func Include(name string, data map[string]interface{}) (string, error) {
	var result string
	files, err := filepath.Glob(name)
	if err != nil {
		return "", err
	}
	// Can't use template.ParseFiles file names can be same
	for _, file := range files {
		vars, err := ValuesFromDirectory(path.Dir(file))
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
		out := fmt.Sprintf(BlockTemplate, file, NormalizeTextBlock(buf.String()))
		result += out
	}
	result = NormalizeTextBlock(result)
	return result, nil
}
