package httpcli

import (
	"net/http"
)

func Status(host string) int {
	resp, err := http.Get(host)

	if err != nil {
		return 404
	}

	return resp.StatusCode
}
