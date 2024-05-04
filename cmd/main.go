package main

import (
	"log"
	"os"

	"github.com/adzsx/dirsgover/internal/httpc"
	"github.com/adzsx/dirsgover/internal/utils"
)

var (
	help string = ` 
Dirsgover - Dir Discoverer in go

Syntax: dirsgover host [options]

Options:
	-h, --help:		Display help message
	-w [wordlist]:		Define a wordlist with dirs to scan
	-f [code1 code2]:	Filter out specific status codes
	-S [code1 code2]:	Show specific status codes (empty for all)
	-r:			Scan for dirs in robots.txt of host
	`
)

func main() {
	args := os.Args

	log.SetFlags(0)

	if len(args) < 3 && !utils.InSclice(args, "--help") && !utils.InSclice(args, "-h") {
		log.Println("Enter --help for help")
		os.Exit(0)
	}

	scan, err := utils.Args(args)
	utils.Err(err)

	if scan.Host == "help" {
		log.Println(help)
		os.Exit(0)
	}

	scan.Host = utils.ValidAddr(scan.Host)

	err = httpc.Scan(scan)
	utils.Err(err)

}
