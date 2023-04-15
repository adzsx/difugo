package main

import (
	"fmt"
	"os"

	"github.com/adzsx/dirsgover/pkg/check"
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

	if !check.ValidHost(scan.Host) {
		fmt.Println("Host Address invalid")
		os.Exit(0)
	}

	status := httpcli.Status(scan.Host)

	fmt.Println(status)

}
