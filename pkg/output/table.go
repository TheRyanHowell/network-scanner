package output

import (
	"fmt"
	"io"
	"strings"
)

// TableWriter writes data in a table format.
type TableWriter struct {
	writer  io.Writer
	headers []string
	widths  []int
}

// NewTableWriter creates a new TableWriter.
func NewTableWriter(writer io.Writer, headers []string) *TableWriter {
	return &TableWriter{
		writer:  writer,
		headers: headers,
		widths:  make([]int, len(headers)),
	}
}

// SetWidths sets fixed column widths.
func (t *TableWriter) SetWidths(widths []int) {
	t.widths = widths
}

// PrintHeader prints the table header.
func (t *TableWriter) PrintHeader() {
	if t.widths[0] == 0 {
		for i, h := range t.headers {
			t.widths[i] = len(h)
		}
	}
	for i, header := range t.headers {
		fmt.Fprintf(t.writer, "%-*s", t.widths[i]+2, header)
	}
	fmt.Fprintln(t.writer)

	for i := range t.headers {
		fmt.Fprintf(t.writer, "%s", strings.Repeat("-", t.widths[i]+2))
	}
	fmt.Fprintln(t.writer)
}

// PrintRow prints a single row.
func (t *TableWriter) PrintRow(row []string) {
	for i, cell := range row {
		if len(cell) > t.widths[i] {
			cell = cell[:t.widths[i]]
		}
		fmt.Fprintf(t.writer, "%-*s", t.widths[i]+2, cell)
	}
	fmt.Fprintln(t.writer)
}
