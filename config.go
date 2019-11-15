package main

import (
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Checks    map[string]*CheckOptions `yaml:"checks"`
	TestFiles []string                 `yaml:"testFiles"`
}

func (c *Config) SetDefaults() *Config {
	if c.Checks == nil {
		c.Checks = make(map[string]*CheckOptions)
	}
	if len(c.TestFiles) == 0 {
		c.TestFiles = []string{
			"*.yaml",
			"*.yml",
		}
	}
	return c
}

func DefaultConfig() *Config {
	return (&Config{}).SetDefaults()
}

func LoadFromFile(filename string) (*Config, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return DefaultConfig(), nil
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
	return c.SetDefaults(), nil
}
