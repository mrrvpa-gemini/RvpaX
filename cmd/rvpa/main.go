package main

import (
	"fmt"
	"os"

	"github.com/mrrvpa-gemini/rvpa/internal/config"
	"github.com/mrrvpa-gemini/rvpa/pkg/banner"
	"github.com/mrrvpa-gemini/rvpa/pkg/colors"
	"github.com/spf13/cobra"
)

var (
	Version = "1.0.0"
	Build   = "dev"
)

var rootCmd = &cobra.Command{
	Use:   "rvpa",
	Short: "Fast Web Scanning & Security Audit Framework",
	Long: `RVPA - Rapid Vulnerability & Penetration Assessment

A comprehensive web scanning tool for security auditing and reconnaissance.
Supports fingerprinting, technology detection, API discovery, and more.

Examples:
  rvpa -u https://example.com
  rvpa -u https://example.com -deep
  rvpa -list targets.txt -json
  rvpa -u https://example.com -html -o report.html`,
	Run: func(cmd *cobra.Command, args []string) {
		if showVersion, _ := cmd.Flags().GetBool("version"); showVersion {
			printVersion()
			return
		}
		if showUpdate, _ := cmd.Flags().GetBool("update"); showUpdate {
			printUpdate()
			return
		}

		banner.PrintBanner()

		url, _ := cmd.Flags().GetString("url")
		listFile, _ := cmd.Flags().GetString("list")

		if url == "" && listFile == "" {
			cmd.Help()
			return
		}

		// Initialize scanner with flags
		initializeScanner(cmd)
	},
}

func init() {
	rootCmd.Flags().StringP("url", "u", "", "Target URL to scan")
	rootCmd.Flags().String("list", "", "File containing list of targets")
	rootCmd.Flags().Bool("deep", false, "Enable deep scanning")
	rootCmd.Flags().Bool("json", false, "Output in JSON format")
	rootCmd.Flags().Bool("html", false, "Generate HTML report")
	rootCmd.Flags().StringP("output", "o", "", "Output file path")
	rootCmd.Flags().IntP("threads", "t", 20, "Number of concurrent threads")
	rootCmd.Flags().Int("timeout", 30, "Request timeout in seconds")
	rootCmd.Flags().String("proxy", "", "HTTP proxy URL")
	rootCmd.Flags().StringP("user-agent", "ua", "", "Custom user agent")
	rootCmd.Flags().Bool("version", false, "Show version information")
	rootCmd.Flags().Bool("update", false, "Update to latest version")
	rootCmd.Flags().BoolP("help", "h", false, "Display help message")
}

func initializeScanner(cmd *cobra.Command) {
	url, _ := cmd.Flags().GetString("url")
	listFile, _ := cmd.Flags().GetString("list")
	deepScan, _ := cmd.Flags().GetBool("deep")
	jsonOutput, _ := cmd.Flags().GetBool("json")
	htmlOutput, _ := cmd.Flags().GetBool("html")
	outputFile, _ := cmd.Flags().GetString("output")
	threads, _ := cmd.Flags().GetInt("threads")
	timeout, _ := cmd.Flags().GetInt("timeout")
	proxy, _ := cmd.Flags().GetString("proxy")
	userAgent, _ := cmd.Flags().GetString("user-agent")

	cfg := &config.ScanConfig{
		URL:       url,
		ListFile:  listFile,
		DeepScan:  deepScan,
		JSON:      jsonOutput,
		HTML:      htmlOutput,
		Output:    outputFile,
		Threads:   threads,
		Timeout:   timeout,
		Proxy:     proxy,
		UserAgent: userAgent,
	}

	fmt.Println(colors.Cyan("[*]") + " Starting scan with config:")
	fmt.Printf("    URL: %s\n", cfg.URL)
	fmt.Printf("    Deep Scan: %v\n", cfg.DeepScan)
	fmt.Printf("    Threads: %d\n", cfg.Threads)
	fmt.Printf("    Timeout: %ds\n", cfg.Timeout)
}

func printVersion() {
	fmt.Printf("RVPA version %s (%s)\n", Version, Build)
	fmt.Println("Go version:", "1.24+")
}

func printUpdate() {
	fmt.Println(colors.Yellow("[!]") + " Update functionality coming soon")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
