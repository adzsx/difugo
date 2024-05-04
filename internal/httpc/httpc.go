package httpc

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/adzsx/dirsgover/internal/utils"
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

func Robots(host string) {
	bleepbloop := host + "/robots.txt"
	status := Status(bleepbloop)
	if status == 200 {
		resp, _ := http.Get(bleepbloop)

		resBody, err := io.ReadAll(resp.Body)
		body := string(resBody)

		var ent string

		for _, entry := range strings.Split(body, "\n") {
			ent = utils.FilterChar(entry, ":", false)
			if len(ent) > 1 {
				if string(ent[1]) == "/" {

					bleep = append(bleep, ent[2:])
					wg.Add(1)
				}
			}

		}
		utils.Err(err)
	} else if status == 404 {
		log.Fatalln("robots.txt does not exist on this host")
	}
}
