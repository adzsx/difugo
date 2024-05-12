package main

import (
	"log"
	"os"

	"github.com/adzsx/axolyn/internal/httpc"
	"github.com/adzsx/axolyn/internal/utils"
)

var (
	help string = ` 
Dirsgover - Dir Discoverer in go

Syntax: dirsgover [options]

Options:
	-h, --help:				Display help message
	-u, --url	[URL]:			Set hosts URL
	-w, --wordlist 	[wordlist]:		Define a wordlist with dirs to scan
	-f, --filter 	[code1 code2]:		Filter out specific status codes
	-s, --show 	[code1 code2]:		Show specific status codes (empty for all)
	-r, --robots:				Scan for dirs in robots.txt of host
	-a, --async 	[threads]:		Use [n] asynchronous threads (default: 32) 
	-v, --verbose:				Verbose mode (print more info)
	--debug:				Enable debug mode (print even more info)
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

	up := httpc.Up(scan.Host)

	if !up {
		log.Fatalln("Host is not up")
	}

	err = httpc.Scan(scan)
	utils.Err(err)

}
