package output

import (
	"bytes"
	"testing"
)

func TestNewCsvWriter(t *testing.T) {
	headers := []string{"h1", "h2", "h3"}
	writer := &bytes.Buffer{}
	csvWriter := NewCsvWriter(writer, headers)

	if csvWriter == nil {
		t.Error("NewCsvWriter should not return nil")
		return
	}

	if csvWriter.writer == nil {
		t.Error("csvWriter.writer should not be nil")
	}

	if len(csvWriter.headers) != len(headers) {
		t.Errorf("expected headers length %d, got %d", len(headers), len(csvWriter.headers))
	}

	for i := range headers {
		if csvWriter.headers[i] != headers[i] {
			t.Errorf("expected header %s, got %s", headers[i], csvWriter.headers[i])
		}
	}
}

func TestCsvWriter_PrintHeader(t *testing.T) {
	headers := []string{"col1", "col2", "col3"}
	writer := &bytes.Buffer{}
	csvWriter := NewCsvWriter(writer, headers)

	csvWriter.PrintHeader()
	csvWriter.writer.Flush()

	expected := "col1,col2,col3\n"
	if writer.String() != expected {
		t.Errorf("expected %q, got %q", expected, writer.String())
	}
}

func TestCsvWriter_PrintRow(t *testing.T) {
	row := []string{"d1", "d2", "d3"}
	writer := &bytes.Buffer{}
	csvWriter := NewCsvWriter(writer, []string{})

	csvWriter.PrintRow(row)

	expected := "d1,d2,d3\n"
	if writer.String() != expected {
		t.Errorf("expected %q, got %q", expected, writer.String())
	}
}

func TestCsvWriter_PrintHeaderAndRow(t *testing.T) {
	headers := []string{"h1", "h2"}
	row := []string{"r1", "r2"}
	writer := &bytes.Buffer{}
	csvWriter := NewCsvWriter(writer, headers)

	csvWriter.PrintHeader()
	csvWriter.PrintRow(row)

	expected := "h1,h2\nr1,r2\n"
	if writer.String() != expected {
		t.Errorf("expected %q, got %q", expected, writer.String())
	}
}
