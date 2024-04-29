package httpc

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

func Up(host string) int {
	fmt.Println("Checking status of host...")

	status := Status(host)

	return status
}
