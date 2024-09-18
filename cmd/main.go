package main

import (
	"log"
	"os"

	"github.com/adzsx/difugo/internal/httpc"
	"github.com/adzsx/difugo/internal/utils"
)

var (
	help string = ` 
difugo - A Directory fuzzer

Usage:
	difugo [options]

Options:
	-h, --help:			Display help message
	
	-u, --url	[string]:	Set hosts URL
	-w, --wordlist 	[file]:		Define a wordlist with dirs to scan
	-r, --robots:			Scan for dirs in robots.txt of host

	-s, --sufix 	[string]:	Set the suffix (example.com/[fuzzed value][suffix])			default: ""
	
	-f, --filter 	[[]int]:	Filter out specific status codes					default: 403, 404
	-c, --code 	[[]int]:	Show specific status codes (if empty, everything is shown)

	-a, --async 	[int]:		Number of asynchronous threads						default: 32
	-l, --level	[int]:		Find directories recursive						default: 1
	-v, --verbose:			Verbose mode
	--debug:			Enable debug mode
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
