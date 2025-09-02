package scanner

import (
	"net"
	"strings"
	"time"
)

// Scanner is the interface for a port scanner.
type Scanner interface {
	Scan(p Port) Port
}

// PortScanner is a concrete implementation of Scanner.
type PortScanner struct {
	Timeout time.Duration
}

// NewPortScanner creates a new PortScanner.
func NewPortScanner(timeout time.Duration) Scanner {
	return &PortScanner{Timeout: timeout}
}

// Scan performs the port scan.
func (ps *PortScanner) Scan(p Port) Port {
	address := p.String()
	conn, err := net.DialTimeout("tcp", address, ps.Timeout)
	if err != nil {
		if strings.Contains(err.Error(), "timeout") {
			p.Status = Timeout
		} else {
			p.Status = Closed
		}
		return p
	}
	defer conn.Close()
	p.Status = Open
	return p
}
