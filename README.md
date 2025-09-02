# Network Scanner

A simple and fast CLI tool built in Go to scan a network for open ports concurrently.

## Features

*   Scan a CIDR range of IP addresses.
*   Scan a list or range of ports.
*   Adjustable timeout for port scans.
*   Output results in a table or CSV format.
*   Filter results to show all, open, or open and timeout ports.

## Installation

To install the network scanner, use `go install`:

```bash
go install github.com/theryanhowell/network-scanner@latest
```

## Usage

```bash
network-scanner [CIDR] [ports] [flags]
```

### Arguments

*   `<CIDR>`: The CIDR range of the network to scan (e.g., `192.168.1.0/24`).
*   `[ports]`: (Optional) A comma-separated list of ports (e.g., `22,80,443`) or a port range (e.g., `1-1024`). Defaults to `1-1024`.

### Flags

*   `--show-all`, `-a`: Show all ports, including closed ones.
*   `--show-open`, `-o`: Only show open ports.
*   `--csv`, `-c`: Output in CSV format.
*   `--timeout`, `-t`: Timeout for each port scan. Defaults to `3s`.

## Examples

Scan the `192.168.1.0/24` network for common web ports:

```bash
network-scanner 192.168.1.0/24 80,443,8080
```

Scan a single host for all ports from 1 to 1024 and output to a CSV file:

```bash
network-scanner 192.168.1.10/32 1-1024 --csv > ports.csv
```

Only show open ports with a 5-second timeout:

```bash
network-scanner 192.168.1.0/24 --show-open -t 5s
```

## Building from Source

To build the network scanner from source, you'll need Go installed.

1.  Clone the repository:
    ```bash
    git clone https://github.com/theryanhowell/network-scanner.git
    ```
2.  Navigate to the project directory:
    ```bash
    cd network-scanner
    ```
3.  Build the application:
    ```bash
    go build
    ```
4.  Run the executable:
    ```bash
    ./network-scanner <CIDR> [ports] [flags]
    ```