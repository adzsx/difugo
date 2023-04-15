package main

import (
	"fmt"
	"os"

	"github.com/adzsx/dirsgover/pkg/format"
	"github.com/adzsx/dirsgover/pkg/httpcli"
)

var (
	help string = ` 
Dirsgover - Dir Discoverer in go

Syntax: dirsgover host [options]

Options:
	-h, --help:					Display help message

	-w wordlist:	Define a wordlist with dirs to scan
	-r				Scan for dirs in robots.txt of host
	`
)

func main() {
	args := os.Args

	scan := format.Args(args)

	if scan.Host == "help" {
		fmt.Println(help)
		os.Exit(0)
	}

	scan = format.Host(scan)

	if scan.Err != "" {
		fmt.Println("Invalid Host address")
		os.Exit(0)
	}

	if httpcli.HostUp(scan.Host) {
		fmt.Println("Host is up")
	} else {
		fmt.Println("Host not reachable")
	}
}
