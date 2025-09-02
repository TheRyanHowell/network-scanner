package output

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTableWriter(t *testing.T) {
	writer := &bytes.Buffer{}
	headers := []string{"h1", "h2"}
	tw := NewTableWriter(writer, headers)

	assert.NotNil(t, tw)
	assert.Equal(t, writer, tw.writer)
	assert.Equal(t, headers, tw.headers)
	assert.NotNil(t, tw.widths)
	assert.Equal(t, len(headers), len(tw.widths))
}

func TestTableWriter_SetWidths(t *testing.T) {
	writer := &bytes.Buffer{}
	headers := []string{"h1", "h2"}
	tw := NewTableWriter(writer, headers)

	widths := []int{10, 20}
	tw.SetWidths(widths)

	assert.Equal(t, widths, tw.widths)
}

func TestTableWriter_PrintHeader_AutoSize(t *testing.T) {
	writer := &bytes.Buffer{}
	headers := []string{"IP", "Port", "Status"}
	tw := NewTableWriter(writer, headers)

	tw.PrintHeader()

	expectedHeader := fmt.Sprintf("%-*s%-*s%-*s\n", 4, "IP", 6, "Port", 8, "Status")
	expectedSeparator := fmt.Sprintf("%s%s%s\n", strings.Repeat("-", 4), strings.Repeat("-", 6), strings.Repeat("-", 8))
	expected := expectedHeader + expectedSeparator
	assert.Equal(t, expected, writer.String())
}

func TestTableWriter_PrintHeader_WithWidths(t *testing.T) {
	writer := &bytes.Buffer{}
	headers := []string{"IP", "Port", "Status"}
	tw := NewTableWriter(writer, headers)
	widths := []int{15, 5, 8}
	tw.SetWidths(widths)

	tw.PrintHeader()

	expectedHeader := fmt.Sprintf("%-*s%-*s%-*s\n", 17, "IP", 7, "Port", 10, "Status")
	expectedSeparator := fmt.Sprintf("%s%s%s\n", strings.Repeat("-", 17), strings.Repeat("-", 7), strings.Repeat("-", 10))
	expected := expectedHeader + expectedSeparator
	assert.Equal(t, expected, writer.String())
}

func TestTableWriter_PrintRow(t *testing.T) {
	writer := &bytes.Buffer{}
	headers := []string{"IP", "Port", "Status"}
	tw := NewTableWriter(writer, headers)
	widths := []int{15, 5, 8}
	tw.SetWidths(widths)

	row := []string{"127.0.0.1", "80", "Open"}
	tw.PrintRow(row)

	expected := fmt.Sprintf("%-*s%-*s%-*s\n", 17, "127.0.0.1", 7, "80", 10, "Open")
	assert.Equal(t, expected, writer.String())
}

func TestTableWriter_FullTable(t *testing.T) {
	writer := &bytes.Buffer{}
	headers := []string{"IP", "Port", "Status"}
	tw := NewTableWriter(writer, headers)
	widths := []int{15, 5, 10}
	tw.SetWidths(widths)

	tw.PrintHeader()
	tw.PrintRow([]string{"127.0.0.1", "80", "Open"})
	tw.PrintRow([]string{"127.0.0.1", "443", "Closed"})
	tw.PrintRow([]string{"255.255.255.255", "12345", "Timed Out"})

	var expected strings.Builder
	expectedHeader := fmt.Sprintf("%-*s%-*s%-*s\n", 17, "IP", 7, "Port", 12, "Status")                                     // Changed from 10 to 12
	expectedSeparator := fmt.Sprintf("%s%s%s\n", strings.Repeat("-", 17), strings.Repeat("-", 7), strings.Repeat("-", 12)) // Changed from 10 to 12
	expected.WriteString(expectedHeader)
	expected.WriteString(expectedSeparator)
	expected.WriteString(fmt.Sprintf("%-*s%-*s%-*s\n", 17, "127.0.0.1", 7, "80", 12, "Open"))               // Changed from 10 to 12
	expected.WriteString(fmt.Sprintf("%-*s%-*s%-*s\n", 17, "127.0.0.1", 7, "443", 12, "Closed"))            // Changed from 10 to 12
	expected.WriteString(fmt.Sprintf("%-*s%-*s%-*s\n", 17, "255.255.255.255", 7, "12345", 12, "Timed Out")) // Changed from 10 to 12

	assert.Equal(t, expected.String(), writer.String())
}
