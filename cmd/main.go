package main

import (
	"log"
	"os"

	"github.com/adzsx/difugo/internal/httpc"
	"github.com/adzsx/difugo/internal/utils"
)

var (
	help string = ` 
Axolyn - Easy Directory fuzzing

Usage:
	axolyn [options]

Options:
	-h, --help:				Display help message
	
	-u, --url	[URL]:			Set hosts URL
	-w, --wordlist 	[wordlist]:		Define a wordlist with dirs to scan
	-r, --robots:				Scan for dirs in robots.txt of host

	-p, --prefix [prefix]:			Set the prefix for the url (example.com[prefix]something)		default: /
	-s, --sufix [suffix]:			Set the suffix (example.com[?]something[=0])				default: ""
	
	
	-f, --filter 	[code1 code2]:		Filter out specific status codes					default: 403, 404
	-c, --code 	[code1 code2]:		Show specific status codes (if empty, everything is shown)

	-a, --async 	[threads]:		Use [n] asynchronous threads (default: 32) 				default: 32
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
