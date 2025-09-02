package iputil

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

// GetIPs returns a list of IPs from a CIDR range.
func GetIPs(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	// remove network and broadcast addresses
	if len(ips) < 2 {
		return ips, nil
	}

	return ips[1 : len(ips)-1], nil
}

// inc increments an IP address to the next one in the network.
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

// ParsePorts parses a comma-separated list of ports or a port range.
func ParsePorts(ports string) ([]int, error) {
	var result []int

	// Handle a list of ports
	if !strings.Contains(ports, "-") {
		parts := strings.Split(ports, ",")
		for _, part := range parts {
			port, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("invalid port: %s", part)
			}
			result = append(result, port)
		}

		return result, nil
	}

	// Handle port ranges
	parts := strings.Split(ports, "-")

	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid port range: %s", ports)
	}

	start, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, fmt.Errorf("invalid start port: %s", parts[0])
	}

	end, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid end port: %s", parts[1])
	}

	if start > end {
		return nil, fmt.Errorf("invalid port range: %s", ports)
	}

	for i := start; i <= end; i++ {
		result = append(result, i)
	}

	return result, nil
}
