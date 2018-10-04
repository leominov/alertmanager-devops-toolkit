package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	yaml "gopkg.in/yaml.v2"
)

const (
	TemplateOption = "missingkey=error"
)

var (
	funcMap = template.FuncMap{
		"indent":    Indent,
		"include":   Include,
		"toYaml":    ToYaml,
		"fromYaml":  FromYaml,
		"normalize": Normalize,
	}
	funcMapBlock = template.FuncMap{
		"indent":    Indent,
		"toYaml":    ToYaml,
		"fromYaml":  FromYaml,
		"normalize": Normalize,
	}
)

func Normalize(str string) string {
	return regexp.MustCompile("[^a-zA-Z0-9]+").ReplaceAllString(str, "")
}

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
		b, err := ioutil.ReadFile(file)
		if err != nil {
			return "", err
		}
		t := template.New(file).Option(TemplateOption).Funcs(funcMapBlock)
		t, err = t.Parse(string(b))
		if err != nil {
			return "", err
		}
		buf := bytes.NewBufferString("")
		if err := t.Execute(buf, data); err != nil {
			return "", err
		}
		out := fmt.Sprintf(BlockTemplate, file, NormalizeTextBlock(buf.String()))
		result += out
	}
	result = NormalizeTextBlock(result)
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
