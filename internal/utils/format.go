package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Input struct {
	Host     string
	Wordlist string
	Robots   bool

	Suffix string

	// Return codes
	// Return codes to print
	StatShow []int
	// Return codes to hide
	StatHide []int

	Workers int
	Level   int
}

var (
	threadLimit int = 2048
)

func Args(args []string) (Input, error) {
	scan := Input{}

	var err error
	vlevel = 0

	for index, element := range args {
		switch element {
		case "-u", "--url":
			scan.Host = args[index+1]
		case "-w", "--wordlist":
			scan.Wordlist = args[index+1]

		case "-r", "--robots":
			scan.Robots = true

		case "-s", "--suffix":
			scan.Suffix = args[index+1]

		case "-c", "--code":
			for i := 1; len(args) > index+i && !strings.Contains(args[index+i], "-"); i++ {
				code, _ := strconv.Atoi(args[index+i])
				scan.StatShow = append(scan.StatShow, code)
			}

			if len(scan.StatShow) == 0 {
				scan.StatShow = []int{}
			}

		case "-f", "--filter":
			for i := 1; len(args) > index+i && !strings.Contains(args[index+i], "-"); i++ {
				code, _ := strconv.Atoi(args[index+i])
				scan.StatHide = append(scan.StatHide, code)

			}
		case "-a", "--async":
			scan.Workers, _ = strconv.Atoi(args[index+1])

		case "-v", "--verbose":
			vlevel = 1

		case "--debug":
			vlevel = 2
		}
	}

	if InSclice(args, "--help") || InSclice(args, "help") || InSclice(args, "-h") {
		scan.Host = "help"
	}

	if scan.Host == "" {
		err = errors.New("host")
	}

	if scan.Workers == 0 {
		scan.Workers = 32
	}

	if scan.Workers > threadLimit {
		scan.Workers = threadLimit
		Verbose(1, "Cant open more than ", threadLimit, " threads")
	}

	if !InSclice(args, "-S") {
		scan.StatHide = append(scan.StatHide, 403, 404)
	}

	Verbose(1, "Scan started")

	var wordlist string = scan.Wordlist
	if scan.Robots {
		wordlist = "/robots.txt"
	}

	fmt.Printf(`
Starting difugo
#########################################################
Host:			%v
Wordlist:		%v
Show-Codes:		%v
Ignore-Codes:		%v
Suffix:			%v
Threads:		%v
#########################################################
`, scan.Host, wordlist, scan.StatShow, scan.StatHide, scan.Suffix, scan.Workers)

	return scan, err
}

func Host(scan Input) Input {
	if strings.Contains(scan.Host, "http") {
		return scan
	}

	scan.Host = "http://" + scan.Host
	return scan
}
