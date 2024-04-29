package utils

import (
	"strings"
)

func Err(err error) {
	if err != nil {
		panic(err)
	}
}

func InSclice(s []string, element string) bool {
	for _, v := range s {
		if element == v {
			return true
		}
	}
	return false
}

func ValidHost(host string) bool {
	return strings.Contains(host, ".")
}
