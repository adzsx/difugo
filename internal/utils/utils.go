package utils

import (
	"log"
)

var (
	vlevel int
)

func FilterChar(str string, char string, before bool) string {
	var final string

	for index, element := range str {
		if before {
			if string(element) != char {
				final += string(element)
			} else {
				return final
			}
		} else {

			if string(element) == char {
				final += str[index+1:]
			}
		}

	}

	return final
}

func Verbose(level int, v ...any) {
	if level <= vlevel {
		log.Print(v...)
	}
}
