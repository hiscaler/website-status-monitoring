package report

type Formatter struct {
	Items []map[string]string
}

func (f *Formatter) Output() *Formatter {
	return f
}

type TxtFormatter struct {
	Formatter
}

func (f *TxtFormatter) Output() *TxtFormatter {
	return f
}

type CsvFormatter struct {
	Formatter
}

func (f *CsvFormatter) Output() *CsvFormatter {
	return f
}
