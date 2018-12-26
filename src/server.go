package main

import (
	"config"
	"report"
	"log"
	"os"
	"bufio"
	"net/http"
	"sync"
	"fmt"
)

var cfg *config.Config
var wg sync.WaitGroup

func init() {
	cfg = &config.Config{
		Debug:        true,
		ReportFormat: report.TxtFormat,
	}
}

func test(url string, chanItem chan report.Item) {
	log.Println("Checking " + url)
	item := report.Item{
		Url:        url,
		Accessible: false,
	}
	resp, err := http.Get(url)
	if err == nil {
		resp.Body.Close()
		item.Accessible = true
	}

	chanItem <- item
}

func main() {
	log.Println("Begin...")
	file, err := os.Open("./data/urls.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var r report.IReport
	var txtR report.TxtReport
	var csvR report.CsvReport
	if cfg.ReportFormat == report.TxtFormat {
		r = &txtR
	} else if cfg.ReportFormat == report.CsvFormat {
		r = &csvR
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	item := make(chan report.Item, 10)
	for scanner.Scan() {
		url := scanner.Text()
		if len(url) <= 0 {
			continue
		}
		wg.Add(1)
		go test(url, item)
	}
	go func(ch chan report.Item, wg *sync.WaitGroup) {
		for {
			select {
			case v := <-ch:
				fmt.Println("Checked result", v)
				r.AddItem(v)
				wg.Done()
			}
		}
	}(item, &wg)
	wg.Wait()

	if err := r.Write(); err != nil {
		log.Println("Save fail, " + err.Error())
	}
	log.Println("Done...")
}
