package format

import (
	"github.com/adzsx/dirsgover/pkg/check"
)

type Scan struct {
	Host     string
	Wordlist string

	Robots bool

	Err string
}

func Args(args []string) Scan {
	scan := Scan{}

	for index, element := range args {
		switch element {
		case "-w":
			scan.Wordlist = args[index+1]

		case "-r":
			scan.Robots = true
		}
	}

	scan.Err = ""

	if scan.Host == "" {
		scan.Err = "host"
	} else if scan.Wordlist == "" {
		scan.Err = "wordlist"
	}

	scan.Host = args[1]

	if check.InSclice(args, "--help") || check.InSclice(args, "help") || check.InSclice(args, "-h") {
		scan.Host = "help"
	}

	return scan
}
