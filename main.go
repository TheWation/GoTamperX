package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/TheWation/GoTamperX/internal/httpclient"
)

func main() {
	printBanner()

	// Parse headers first and remove them from os.Args
	var headers []string
	var filteredArgs []string
	for i, arg := range os.Args {
		if arg == "-H" && i+1 < len(os.Args) {
			headers = append(headers, os.Args[i+1])
			i++ // Skip the next argument as it's the header value
		} else if arg != "-H" {
			filteredArgs = append(filteredArgs, arg)
		}
	}
	
	// Replace os.Args with filtered arguments
	os.Args = filteredArgs

	var targetURL string
	var timeout int
	var concurrency int
	var proxy string
	var randomAgent bool
	flag.StringVar(&targetURL, "u", "", "Target URL to test")
	flag.IntVar(&timeout, "t", 10, "Timeout in seconds (default: 10)")
	flag.IntVar(&concurrency, "c", 8, "Number of concurrent requests (default: 8)")
	flag.StringVar(&proxy, "p", "", "Proxy URL (e.g., http://proxy:8080)")
	flag.BoolVar(&randomAgent, "random-agent", false, "Use random User-Agent for each request")
	flag.Parse()

	if targetURL == "" {
		fmt.Println("[-] URL not provided")
		fmt.Println("Usage: tamperx -u <url> [-t timeout] [-c concurrency] [-p proxy] [--random-agent] [-H header]")
		os.Exit(1)
	}

	client := httpclient.NewClient(timeout, proxy, randomAgent, headers)
	httpMethods := []string{"GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "TRACE", "PATCH"}

	fmt.Printf("\n[+] Target Url: %s\n", targetURL)
	if proxy != "" {
		fmt.Printf("[+] Using Proxy: %s\n", proxy)
	}
	if randomAgent {
		fmt.Printf("[+] Random User-Agent: Enabled\n")
	}
	if len(headers) > 0 {
		fmt.Printf("[+] Custom Headers: %d\n", len(headers))
	}
	fmt.Printf("[+] Concurrency: %d\n", concurrency)
	fmt.Printf("[+] Timeout: %ds\n\n", timeout)

	fmt.Printf("%-10s %-10s %-11s\n", "Method", "Status", "Content")
	fmt.Println(strings.Repeat("-", 32))

	results := client.TestMethods(targetURL, httpMethods, concurrency)

	for _, result := range results {
		if result.Error != nil {
			fmt.Printf("%-10s %-10s %-11s (%s)\n", result.Method, "ERROR", "0", result.Error.Error())
		} else {
			fmt.Printf("%-10s %-10d %-11d\n", result.Method, result.StatusCode, result.ContentLen)
		}
	}
}

func printBanner() {
	fmt.Println("  __     __     ______     ______   __     ______     __   __    ")
	fmt.Println(" /\\ \\  _ \\ \\   /\\  __ \\   /\\__  _\\ /\\ \\   /\\  __ \\   /\\ \"-.\\ \\   ")
	fmt.Println(" \\ \\ \\/ \".\\ \\  \\ \\  __ \\  \\/_/\\ \\/ \\ \\ \\  \\ \\ \\/\\ \\  \\ \\ \\-.  \\  ")
	fmt.Println("  \\ \\__/\".~\\_\\  \\ \\_\\ \\_\\    \\ \\_\\  \\ \\_\\  \\ \\_____\\  \\ \\_\\\"\\_ \\ ")
	fmt.Println("   \\/_/   \\/_/   \\/_/\\/_/     \\/_/   \\/_/   \\/_____/   \\/_/ \\/_/ ")
	fmt.Println("        TamperX V2.0: Verb Tampering Vulnerability Checker")
}
