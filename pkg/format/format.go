package format

type Scan struct {
	mode     string
	host     string
	wordlist string

	robots bool

	err string
}

func Args(args []string) Scan {
	scan := Scan{}

	scan.mode = args[1]

	for index, element := range args {
		switch element {
		case "-h":
			scan.host = args[index+1]

		case "-w":
			scan.wordlist = args[index+1]

		case "-r":
			scan.robots = true
		}
	}

	scan.err = ""

	if scan.mode == "" {
		scan.err = "mode"
	} else if scan.host == "" {
		scan.err = "host"
	} else if scan.wordlist == "" {
		scan.err = "wordlist"
	}

	return scan
}
