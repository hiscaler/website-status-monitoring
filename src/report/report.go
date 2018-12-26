package report

import (
	"fmt"
	"time"
	"io/ioutil"
	"log"
	"os"
	"encoding/csv"
)

const (
	TxtFormat = "txt"
	CsvFormat = "csv"
)

type IReport interface {
	Write() error
}

type Item struct {
	Url        string `title:"网址"`
	Datetime   string `title:"时间"`
	Accessible bool   `title:"状态"`
	Status     string `title:"状态"`
}

type Report struct {
	formatter string
	items     []Item
}

func (r *Report) SetFormatter(formatter string) {
	r.formatter = formatter
}

func (r *Report) Filename() string {
	ext := ".txt"
	switch r.formatter {
	case TxtFormat:
		ext = ".txt"

	case CsvFormat:
		ext = ".csv"

	default:
		ext = ".txt"
	}

	now := time.Now()

	dir := "./data/reports/" + now.Format("20060102")
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(dir, 0644)
		}
	}

	return dir + "/" + now.Format("2006-01-02 15.04.05") + ext
}

func (r *Report) AddItem(item Item) {
	item.Datetime = time.Now().Format("2006-01-02 15:04:05")
	if item.Accessible {
		item.Status = "√"
	} else {
		item.Status = "×"
	}
	r.items = append(r.items, item)
}

func (t *Report) Write() error {
	fmt.Println("Report Write")
	for k, v := range t.items {
		fmt.Println(k)
		fmt.Println(v)
	}

	return nil
}

func NewReport() *Report {
	return &Report{}
}

type TxtReport struct {
	Report
}

func (r *TxtReport) Write() error {
	log.Println("Start save text report...")
	rows := []byte{}
	rows = append(rows, []byte(fmt.Sprintf("%s\t%-60s\t%-20s\t%s\n", "序号", "URL", "时间", "是否可访问"))...)
	for k, v := range r.items {
		row := fmt.Sprintf("%-4d\t%-60s\t%-20s\t%s\n", k+1, v.Url, v.Datetime, v.Status)
		rows = append(rows, []byte(row)...)
	}
	err := ioutil.WriteFile(r.Filename(), rows, 0644)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("TxtReport write success.")

	return nil
}

type CsvReport struct {
	Report
}

func (r *CsvReport) Write() error {
	csvFile, err := os.Create(r.Filename())
	if err != nil {
		return err
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()
	for _, item := range r.items {
		d := []string{
			item.Url,
			item.Datetime,
			item.Status,
		}
		if err := csvWriter.Write(d); err != nil {
			log.Println("Error writing record to csv:", err)
			return err
		}
	}
	return nil
}
