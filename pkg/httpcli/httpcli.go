package httpcli

import (
	"fmt"
	"net/http"
)

func Status(host string) int {
	resp, err := http.Get(host)

	if err != nil {
		return 404
	}

	return resp.StatusCode
}

func HostUp(host string) bool {
	fmt.Println("Checking status of host...")

	status := Status(host)

	if status > 99 && status < 400 {
		return true
	}
	return false
}
