package scanner

import (
	"net"
	"testing"
	"time"
)

func TestPortScanner_Open(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:1337")
	if err != nil {
		t.Fatalf("failed to create listener: %v", err)
	}
	defer listener.Close()

	p := Port{
		Host: "127.0.0.1",
		Port: 1337,
	}

	scanner := NewPortScanner(3 * time.Second)
	result := scanner.Scan(p)

	if result.Status != Open {
		t.Errorf("expected status Open, got %v", result.Status)
	}
}

func TestPortScanner_Closed(t *testing.T) {
	p := Port{
		Host: "127.0.0.1",
		Port: 1,
	}

	scanner := NewPortScanner(3 * time.Second)
	result := scanner.Scan(p)

	if result.Status != Closed {
		t.Errorf("expected status Closed, got %v", result.Status)
	}
}

func TestPortScanner_Timeout(t *testing.T) {
	p := Port{
		Host: "192.0.2.1",
		Port: 80,
	}

	scanner := NewPortScanner(1 * time.Millisecond)
	result := scanner.Scan(p)

	if result.Status != Timeout {
		t.Errorf("expected status Timeout, got %v", result.Status)
	}
}
