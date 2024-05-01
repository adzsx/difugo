package utils

import (
	"net"
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

func ValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

func ValidAddr(url string) string {
	if !ValidIP(url) {
		if !strings.Contains(url, "https://") && !strings.Contains(url, "http://") {
			return "https://" + url
		}
	} else {
		return "http://" + url
	}

	return url
}
