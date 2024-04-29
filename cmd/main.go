package main

import (
	"log"
	"os"

	"github.com/adzsx/dirsgover/internal/check"
	"github.com/adzsx/dirsgover/internal/format"
	"github.com/adzsx/dirsgover/internal/httpcli"
)

var (
	help string = ` 
Dirsgover - Dir Discoverer in go

Syntax: dirsgover host [options]

Options:
	-h, --help:	Display help message
	-w wordlist:	Define a wordlist with dirs to scan
	-r:		Scan for dirs in robots.txt of host
	`
)

func main() {
	args := os.Args

	log.SetFlags(0)

	if len(args) < 3 && !check.InSclice(args, "--help") && !check.InSclice(args, "-h") {
		log.Fatalln("Enter --help for help")
	}

	scan := format.Args(args)

	if scan.Host == "help" {
		log.Println(help)
		os.Exit(0)
	}

	scan = format.Host(scan)

	if scan.Err != "" {
		log.Println("Invalid Host address")
		os.Exit(0)
	}

	if httpcli.HostUp(scan.Host) {
		log.Println("Host is up")
	} else {
		log.Fatalln("Host not reachable")
	}
}
