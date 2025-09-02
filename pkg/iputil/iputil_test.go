package iputil

import (
	"net"
	"reflect"
	"testing"
)

func TestGetIPs(t *testing.T) {
	testCases := []struct {
		name     string
		cidr     string
		expected []string
		hasError bool
	}{
		{
			name:     "valid /30 CIDR",
			cidr:     "192.168.1.0/30",
			expected: []string{"192.168.1.1", "192.168.1.2"},
			hasError: false,
		},
		{
			name:     "valid /29 CIDR",
			cidr:     "192.168.1.0/29",
			expected: []string{"192.168.1.1", "192.168.1.2", "192.168.1.3", "192.168.1.4", "192.168.1.5", "192.168.1.6"},
			hasError: false,
		},
		{
			name:     "CIDR with no usable IPs /31",
			cidr:     "192.168.1.0/31",
			expected: []string{},
			hasError: false,
		},
		{
			name:     "CIDR with one IP /32",
			cidr:     "192.168.1.1/32",
			expected: []string{"192.168.1.1"},
			hasError: false,
		},
		{
			name:     "invalid CIDR",
			cidr:     "invalid-cidr",
			expected: nil,
			hasError: true,
		},
		{
			name:     "IPv6 CIDR /126",
			cidr:     "2001:db8::/126",
			expected: []string{"2001:db8::1", "2001:db8::2"},
			hasError: false,
		},
		{
			name:     "IPv6 CIDR /127",
			cidr:     "2001:db8::/127",
			expected: []string{},
			hasError: false,
		},
		{
			name:     "IPv6 CIDR /128",
			cidr:     "2001:db8::1/128",
			expected: []string{"2001:db8::1"},
			hasError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ips, err := GetIPs(tc.cidr)
			if (err != nil) != tc.hasError {
				t.Errorf("GetIPs() error = %v, wantErr %v", err, tc.hasError)
				return
			}
			if !reflect.DeepEqual(ips, tc.expected) {
				t.Errorf("GetIPs() = %v, want %v", ips, tc.expected)
			}
		})
	}
}

func TestInc(t *testing.T) {
	testCases := []struct {
		name     string
		input    net.IP
		expected net.IP
	}{
		{
			name:     "simple increment",
			input:    net.ParseIP("192.168.1.1").To4(),
			expected: net.ParseIP("192.168.1.2").To4(),
		},
		{
			name:     "byte rollover",
			input:    net.ParseIP("192.168.1.255").To4(),
			expected: net.ParseIP("192.168.2.0").To4(),
		},
		{
			name:     "multiple byte rollover",
			input:    net.ParseIP("192.168.255.255").To4(),
			expected: net.ParseIP("192.169.0.0").To4(),
		},
		{
			name:     "all bytes rollover",
			input:    net.ParseIP("255.255.255.255").To4(),
			expected: net.ParseIP("0.0.0.0").To4(),
		},
		{
			name:     "simple IPv6 increment",
			input:    net.ParseIP("2001:db8::1"),
			expected: net.ParseIP("2001:db8::2"),
		},
		{
			name:     "IPv6 rollover",
			input:    net.ParseIP("2001:db8::ffff"),
			expected: net.ParseIP("2001:db8::1:0"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			inc(tc.input)
			if !reflect.DeepEqual(tc.input, tc.expected) {
				t.Errorf("inc() got %v, want %v", tc.input, tc.expected)
			}
		})
	}
}

func TestParsePorts(t *testing.T) {
	testCases := []struct {
		name     string
		ports    string
		expected []int
		hasError bool
	}{
		{
			name:     "single port",
			ports:    "80",
			expected: []int{80},
			hasError: false,
		},
		{
			name:     "comma-separated ports",
			ports:    "80,443,8080",
			expected: []int{80, 443, 8080},
			hasError: false,
		},
		{
			name:     "port range",
			ports:    "80-82",
			expected: []int{80, 81, 82},
			hasError: false,
		},
		{
			name:     "single port in range",
			ports:    "80-80",
			expected: []int{80},
			hasError: false,
		},
		{
			name:     "invalid port in list",
			ports:    "80,abc,443",
			expected: nil,
			hasError: true,
		},
		{
			name:     "invalid port range format",
			ports:    "80-90-100",
			expected: nil,
			hasError: true,
		},
		{
			name:     "invalid start port in range",
			ports:    "abc-100",
			expected: nil,
			hasError: true,
		},
		{
			name:     "invalid end port in range",
			ports:    "80-abc",
			expected: nil,
			hasError: true,
		},
		{
			name:     "start port greater than end port",
			ports:    "100-80",
			expected: nil,
			hasError: true,
		},
		{
			name:     "mixed format (unsupported)",
			ports:    "80,443-8080",
			expected: nil,
			hasError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ports, err := ParsePorts(tc.ports)
			if (err != nil) != tc.hasError {
				t.Errorf("ParsePorts() error = %v, wantErr %v", err, tc.hasError)
				return
			}
			if !reflect.DeepEqual(ports, tc.expected) {
				t.Errorf("ParsePorts() = %v, want %v", ports, tc.expected)
			}
		})
	}
}
