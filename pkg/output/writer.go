package output

// OutputWriter is an interface for writing data in different formats.
type OutputWriter interface {
	PrintHeader()
	PrintRow(row []string)
}
