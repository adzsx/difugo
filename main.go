package main

import (
	"fmt"
	"os"

	"github.com/adzsx/dirsgover/pkg/format"
)

var (
	start string = ``

	help string = `
Dirsgover - Dir Discoverer in go

Syntax: dirsgover mode [options]

option:
	-h host:		Define the host to scan (required)	
	--help:			Display help message		

	-w wordlis:		Define a wordlist with dirs to scan
	-r 				Scan for dirs in robots.txt of host
	`
)

func main() {
	args := os.Args

	scan := format.Args(args)

	fmt.Println(scan)
		
}
