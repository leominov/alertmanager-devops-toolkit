package main

import (
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

var (
	defaultConfig = &Config{
		Checks: make(map[string]*CheckOptions),
	}
)

type Config struct {
	Checks map[string]*CheckOptions `yaml:"checks"`
}

func LoadFromFile(filename string) (*Config, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return defaultConfig, nil
	}
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	c := &Config{}
	err = yaml.Unmarshal(b, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
