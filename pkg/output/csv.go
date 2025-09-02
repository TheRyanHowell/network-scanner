package output

import (
	"encoding/csv"
	"io"
)

// CsvWriter writes data in a CSV format.
type CsvWriter struct {
	writer  *csv.Writer
	headers []string
}

// NewCsvWriter creates a new CsvWriter.
func NewCsvWriter(writer io.Writer, headers []string) *CsvWriter {
	return &CsvWriter{
		writer:  csv.NewWriter(writer),
		headers: headers,
	}
}

// PrintHeader writes the headers to the CSV file.
func (c *CsvWriter) PrintHeader() {
	c.writer.Write(c.headers)
}

// PrintRow writes a single row to the CSV file.
func (c *CsvWriter) PrintRow(row []string) {
	c.writer.Write(row)
	c.writer.Flush()
}
