package main

import "testing"

func TestLoadFromFile(t *testing.T) {
	_, err := LoadFromFile("conf/config/config_not_found.yaml")
	if err != nil {
		t.Error(err)
	}
	_, err = LoadFromFile("conf/config/config_invalid.yaml")
	if err == nil {
		t.Error("Must be an error, but got nil")
	}
	_, err = LoadFromFile("conf/config/config_valid.yaml")
	if err != nil {
		t.Error(err)
	}
}
