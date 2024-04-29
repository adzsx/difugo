package utils

import (
	"errors"
	"strings"
)

type Input struct {
	Host     string
	Wordlist string
	Robots   bool
}

func Args(args []string) (Input, error) {
	scan := Input{}
	var err error

	for index, element := range args {
		switch element {
		case "-w":
			scan.Wordlist = args[index+1]

		case "-r":
			scan.Robots = true
		}
	}

	scan.Host = args[1]

	if InSclice(args, "--help") || InSclice(args, "help") || InSclice(args, "-h") {
		scan.Host = "help"
	}

	if scan.Host == "" {
		err = errors.New("host")
	}

	return scan, err
}

func Host(scan Input) Input {
	if strings.Contains(scan.Host, "http") {
		return scan
	}

	scan.Host = "http://" + scan.Host
	return scan
}
