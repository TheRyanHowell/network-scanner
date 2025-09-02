package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/theryanhowell/network-scanner/pkg/iputil"
	"github.com/theryanhowell/network-scanner/pkg/output"
	"github.com/theryanhowell/network-scanner/pkg/scanner"

	"github.com/spf13/cobra"
)

var (
	showAll  bool
	showOpen bool
	csv      bool
	timeout  time.Duration
)

var rootCmd = &cobra.Command{
	Use:   "network-scanner [CIDR] [ports]",
	Short: "A simple network scanner",
	Long:  `A simple CLI tool built in Go to scan a network for open ports.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ports := "1-1024"
		if len(args) > 1 {
			ports = args[1]
		}

		cidr := args[0]
		ipList, err := iputil.GetIPs(cidr)
		if err != nil {
			fmt.Println("Error getting IPs from CIDR:", err)
			os.Exit(1)
		}

		portList, err := iputil.ParsePorts(ports)
		if err != nil {
			fmt.Println("Error parsing ports:", err)
			os.Exit(1)
		}

		var portsToScan []scanner.Port
		for _, ip := range ipList {
			for _, port := range portList {
				portsToScan = append(portsToScan, scanner.Port{Host: ip, Port: port})
			}
		}

		portScanner := scanner.NewPortScanner(timeout)
		worker := scanner.NewWorker(portScanner, portsToScan)
		scanResults := worker.Run()

		headers := []string{"IP Address", "Port", "Status"}
		var writer output.OutputWriter
		if csv {
			writer = output.NewCsvWriter(os.Stdout, headers)
		} else {
			tableWriter := output.NewTableWriter(os.Stdout, headers)
			tableWriter.SetWidths([]int{25, 10, 10})
			writer = tableWriter
		}

		writer.PrintHeader()

		for port := range scanResults {
			status := port.Status
			if showAll || (status == scanner.Open) || (!showOpen && status == scanner.Timeout) {
				writer.PrintRow([]string{port.Host, fmt.Sprintf("%d", port.Port), port.Status.String()})
			}
		}
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&showAll, "show-all", "a", false, "Show all ports, including closed ones")
	rootCmd.Flags().BoolVarP(&showOpen, "show-open", "o", false, "Only show open ports")
	rootCmd.Flags().BoolVarP(&csv, "csv", "c", false, "Output in CSV format")
	rootCmd.Flags().DurationVarP(&timeout, "timeout", "t", 3*time.Second, "Timeout for each port scan")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
