package http

import (
	"fmt"
	"net/http"
)

func Status(host string) {
	resp, err := http.Get(host)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.StatusCode)
}
