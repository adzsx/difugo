package format

import "fmt"

type Actions struct {
	host     string
	wordlist string

	robots bool
}

func Args(args []string) {
	filtered := Actions{args[1], args[1], false}

	fmt.Println(filtered)
}
