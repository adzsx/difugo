package format

import (
	"strings"

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

	scan.Host = args[1]

	if check.InSclice(args, "--help") || check.InSclice(args, "help") || check.InSclice(args, "-h") {
		scan.Host = "help"
	}

	if scan.Host == "" {
		scan.Err = "host"
	}
	return scan
}

func Host(scan Scan) Scan {
	if strings.Contains(scan.Host, "http") {
		return scan
	}

	scan.Host = "http://" + scan.Host
	return scan
}
