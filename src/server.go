package main

import (
	"config"
	"report"
	"log"
)

var cfg *config.Config

func init() {
	report := &report.Report{
		Enable:    true,
		Formatter: report.CsvFormat,
	}
	cfg = &config.Config{
		Debug:  true,
		Report: *report,
	}
}

func main() {
	log.Println("Begin...")
	switch cfg.Report.Formatter {
	case report.TxtFormat:
		formatter := &report.TxtFormatter{}
		formatter.Output()

	case report.CsvFormat:
		formatter := &report.CsvFormatter{}
		formatter.Output()
	}
	log.Println("Done...")
}
