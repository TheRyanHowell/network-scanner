package scanner

import "testing"

func TestStatus_String(t *testing.T) {
	tests := []struct {
		name   string
		status Status
		want   string
	}{
		{
			name:   "Open",
			status: Open,
			want:   "Open",
		},
		{
			name:   "Closed",
			status: Closed,
			want:   "Closed",
		},
		{
			name:   "Timeout",
			status: Timeout,
			want:   "Timed Out",
		},
		{
			name:   "Unknown",
			status: Status(99),
			want:   "Unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.status.String(); got != tt.want {
				t.Errorf("Status.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPort_String(t *testing.T) {
	p := Port{
		Host: "localhost",
		Port: 8080,
	}

	want := "localhost:8080"

	if got := p.String(); got != want {
		t.Errorf("Port.String() = %v, want %v", got, want)
	}
}
