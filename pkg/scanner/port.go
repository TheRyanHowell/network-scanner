package scanner

import "fmt"

// Status represents the status of a port.
type Status int

const (
	Open Status = iota
	Closed
	Timeout
)

func (s Status) String() string {
	switch s {
	case Open:
		return "Open"
	case Closed:
		return "Closed"
	case Timeout:
		return "Timed Out"
	default:
		return "Unknown"
	}
}

// Port represents a single port on a single host.
type Port struct {
	Host   string
	Port   int
	Status Status
}

func (p Port) String() string {
	return fmt.Sprintf("%s:%d", p.Host, p.Port)
}
