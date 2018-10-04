package main

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

func ValuesFromDirectory(dir string) (map[string]interface{}, error) {
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

func ValuesFromEnviron() map[string]string {
	envs := os.Environ()
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
