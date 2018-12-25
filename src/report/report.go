package report

import (
	"fmt"
	"time"
	"io/ioutil"
	"log"
)

const (
	TxtFormat = "txt"
	CsvFormat = "csv"
)

type IReport interface {
	Write() error
}

type Item struct {
	Url        string    `title:"网址"`
	Datetime   time.Time `title:"时间"`
	Accessible bool      `title:"状态"`
}

type Report struct {
	formatter string
	items     []Item
}

func (r *Report) SetFormatter(formatter string) {
	r.formatter = formatter
}

func (r *Report) AddItem(item Item) {
	item.Datetime = time.Now()
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
	rows := []byte{}
	rows = append(rows, []byte(fmt.Sprintf("%s\t%-20s\t%-20s\t%s\n", "序号", "URL", "时间", "是否可访问"))...)
	for k, v := range r.items {
		row := fmt.Sprintf("%-4d\t%-20s\t%-20s\t%v\n", k+1, v.Url, v.Datetime.Format("2006-01-02 15:04:05"), v.Accessible)
		rows = append(rows, []byte(row)...)
	}
	filename := "./runtime/reports/" + time.Now().Format("20060102150405") + ".txt"
	err := ioutil.WriteFile(filename, rows, 0644)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("TxtReport write success.")

	return nil
}
