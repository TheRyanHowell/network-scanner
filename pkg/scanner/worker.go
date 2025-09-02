package scanner

import "sync"

// Worker manages the concurrent scanning of ports.
type Worker struct {
	scanner Scanner
	ports   []Port
}

// NewWorker creates a new Worker.
func NewWorker(scanner Scanner, ports []Port) *Worker {
	return &Worker{
		scanner: scanner,
		ports:   ports,
	}
}

// Run starts the concurrent scanning and returns a channel of results.
func (w *Worker) Run() <-chan Port {
	resultsChan := make(chan Port)
	var wg sync.WaitGroup

	for _, p := range w.ports {
		wg.Add(1)
		go func(port Port) {
			defer wg.Done()
			resultsChan <- w.scanner.Scan(port)
		}(p)
	}

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	return resultsChan
}
