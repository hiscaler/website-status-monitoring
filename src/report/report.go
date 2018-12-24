package report

const (
	TxtFormat = "txt"
	CsvFormat = "csv"
)

type Report struct {
	Enable    bool
	Formatter string
}
