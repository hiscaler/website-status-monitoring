package config

import (
	"report"
	"io/ioutil"
	"os"
	"log"
	"encoding/json"
)

type Config struct {
	Debug        bool
	ReportFormat string
}

func NewConfig() *Config {
	config := &Config{
		Debug:        true,
		ReportFormat: report.TxtFormat,
	}

	file, err := os.Open("./config/config.json")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	s, err := ioutil.ReadAll(file)

	if err := json.Unmarshal(s, &config); err != nil {
		log.Panic(err)
	}

	return config
}
