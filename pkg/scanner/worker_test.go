package scanner

import (
	"reflect"
	"testing"
)

// MockScanner is a mock implementation of the Scanner interface.
type MockScanner struct {
	ScanFunc func(p Port) Port
}

// Scan implements the Scanner interface for MockScanner.
func (m *MockScanner) Scan(p Port) Port {
	if m.ScanFunc != nil {
		return m.ScanFunc(p)
	}

	p.Status = Open
	return p
}

func TestNewWorker(t *testing.T) {
	ports := []Port{{Port: 80}, {Port: 443}}
	mockScanner := &MockScanner{}
	worker := NewWorker(mockScanner, ports)

	if worker.scanner != mockScanner {
		t.Errorf("expected scanner to be %v, got %v", mockScanner, worker.scanner)
	}

	if !reflect.DeepEqual(worker.ports, ports) {
		t.Errorf("expected ports to be %v, got %v", ports, worker.ports)
	}
}

func TestWorker_Run(t *testing.T) {
	ports := []Port{{Port: 80}, {Port: 443}, {Port: 8080}}

	mockScanner := &MockScanner{}

	worker := NewWorker(mockScanner, ports)
	resultsChan := worker.Run()

	results := make(map[int]Status)
	for p := range resultsChan {
		results[p.Port] = p.Status
	}

	if len(results) != len(ports) {
		t.Errorf("expected %d results, got %d", len(ports), len(results))
	}

	for _, p := range ports {
		status, ok := results[p.Port]
		if !ok {
			t.Errorf("expected result for port %d, but none found", p.Port)
		}
		if status != Open {
			t.Errorf("expected port %d to be open, but it was %s", p.Port, status)
		}
	}
}

func TestWorker_Run_Concurrent(t *testing.T) {
	ports := []Port{{Port: 80}, {Port: 443}, {Port: 8080}, {Port: 22}, {Port: 21}}

	scannedPorts := make(chan int, len(ports))

	mockScanner := &MockScanner{
		ScanFunc: func(p Port) Port {
			scannedPorts <- p.Port
			p.Status = Open
			return p
		},
	}

	worker := NewWorker(mockScanner, ports)
	resultsChan := worker.Run()

	for range resultsChan {
	}

	close(scannedPorts)

	scannedPortNumbers := make(map[int]bool)
	for p := range scannedPorts {
		scannedPortNumbers[p] = true
	}

	if len(scannedPortNumbers) != len(ports) {
		t.Errorf("expected %d ports to be scanned, but got %d", len(ports), len(scannedPortNumbers))
	}

	for _, p := range ports {
		if !scannedPortNumbers[p.Port] {
			t.Errorf("port %d was not scanned", p.Port)
		}
	}
}
