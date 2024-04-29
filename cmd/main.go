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
	-h, --help:	Display help message
	-w wordlist:	Define a wordlist with dirs to scan
	-r:		Scan for dirs in robots.txt of host
	`
)

func main() {
	args := os.Args

	log.SetFlags(0)

	if len(args) < 3 && !utils.InSclice(args, "--help") && !utils.InSclice(args, "-h") {
		log.Fatalln("Enter --help for help")
	}

	scan, err := utils.Args(args)

	utils.Err(err)

	if scan.Host == "help" {
		log.Println(help)
		os.Exit(0)
	}

	scan = utils.Host(scan)

	log.Println(httpc.Up(scan.Host))
}
