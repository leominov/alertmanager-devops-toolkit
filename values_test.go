package main

import (
	"os"
	"testing"
)

func TestValuesFromEnviron(t *testing.T) {
	os.Setenv("FOOBAR", "FOOBAR")
	values := ValuesFromEnviron()
	v, ok := values["FOOBAR"]
	if !ok {
		t.Error("ValuesFromEnviron(FOOBAR) == empty")
	}
	if v != "FOOBAR" {
		t.Error("ValuesFromEnviron() != FOOBAR")
	}
	os.Unsetenv("FOOBAR")
}

func TestValuesFromDirectory(t *testing.T) {
	_, err := ValuesFromDirectory("conf")
	if err == nil {
		t.Errorf("Must be an error, but got nil")
	}
	values, err := ValuesFromDirectory("conf/values1")
	if err != nil {
		t.Error(err)
	}
	v, ok := values["Pref"]
	if !ok {
		t.Error("ValuesFromDirectory(Pref) == empty")
	}
	if v.(string) != "group1" {
		t.Error("ValuesFromDirectory() != group1")
	}
	_, err = ValuesFromDirectory("conf/values2")
	if err == nil {
		t.Errorf("Must be an error, but got nil")
	}
}
