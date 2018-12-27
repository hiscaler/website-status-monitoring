package report

import (
	"fmt"
	"time"
	"io/ioutil"
	"log"
	"os"
	"encoding/csv"
	"reflect"
)

const (
	TxtFormat = "txt"
	CsvFormat = "csv"
)

type IReport interface {
	Write() error
	AddItem(item Item)
}

type Item struct {
	Url        string `网址`
	Datetime   string `时间`
	Accessible bool   `状态`
	Status     string `状态`
}

func GetItemTitles() map[string]string {
	titles := map[string]string{}
	t := reflect.TypeOf(Item{})
	n := t.NumField()
	for i := 0; i < n; i++ {
		f := t.Field(i)
		titles[f.Name] = string(f.Tag)
	}

	return titles
}

type Report struct {
	IReport
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
	r.SetFormatter(TxtFormat)
	titles := GetItemTitles()
	rows := []byte{}
	rows = append(rows, []byte(fmt.Sprintf("%s\t%-60s\t%-20s\t%s\n", "序号", titles["Url"], titles["Datetime"], titles["Status"]))...)
	for k, v := range r.items {
		row := fmt.Sprintf("%-4d\t%-60s\t%-20s\t%s\n", k+1, v.Url, v.Datetime, v.Status)
		rows = append(rows, []byte(row)...)
	}
	err := ioutil.WriteFile(r.Filename(), rows, 0644)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

type CsvReport struct {
	Report
}

func (r *CsvReport) Write() error {
	r.SetFormatter(CsvFormat)
	csvFile, err := os.Create(r.Filename())
	if err != nil {
		return err
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()
	titles := GetItemTitles()
	csvWriter.Write([]string{
		"序号",
		titles["Url"],
		titles["Datetime"],
		titles["Status"],
	})
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
