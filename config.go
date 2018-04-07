package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Address      string
	ReadTimeout  int
	WriteTimeout int
	Static       string
}

var config *Config

func NewConfig() *Config {
	if config != nil {
		return config
	}
	f, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	config = &Config{}
	err = json.Unmarshal(f, config)
	if err != nil {
		panic(err)
	}
	return config
}
