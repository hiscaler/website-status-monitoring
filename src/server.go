package main

import (
	"config"
	"report"
	"log"
	"os"
	"bufio"
)

var cfg *config.Config

func init() {
	cfg = &config.Config{
		Debug:        true,
		ReportFormat: report.TxtFormat,
	}
}

func main() {
	log.Println("Begin...")
	file, err := os.Open("./data/urls.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	r := report.TxtReport{}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		item := report.Item{
			Url:        scanner.Text(),
			Accessible: false,
		}
		r.AddItem(item)
	}

	if err := r.Write(); err != nil {
		log.Println("Save fail, " + err.Error())
	}
	log.Println("Done...")
}
