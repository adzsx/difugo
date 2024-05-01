package httpc

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/adzsx/dirsgover/internal/utils"
)

var (
	scan  utils.Input
	bleep []string
)

func Scan(input utils.Input) error {
	scan = input

	if scan.Robots {
		Robots(scan.Host)
		log.Println(bleep)
		for _, entry := range bleep {
			resp, _ := http.Get(scan.Host + "/" + entry)
			if resp.StatusCode != 404 {
				fmt.Printf("[%v] %v/%v\n", resp.StatusCode, scan.Host, entry)
			}
		}
	} else {
		file, err := os.Open(scan.Wordlist)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		// optionally, resize scanner's capacity for lines over 64K, see next example
		for scanner.Scan() {
			resp, _ := http.Get(scan.Host + "/" + scanner.Text())

			if resp.StatusCode != 404 {
				fmt.Printf("[%v] %v/%v\n", resp.StatusCode, scan.Host, scanner.Text())
			}
		}
	}

	return nil
}

func Robots(host string) {
	bleepbloop := host + "/robots.txt"
	log.Println(bleepbloop)
	status := Status(bleepbloop)
	log.Println(status)
	if status == 200 {
		resp, _ := http.Get(bleepbloop)

		resBody, err := io.ReadAll(resp.Body)
		body := string(resBody)

		var ent string

		for _, entry := range strings.Split(body, "\n") {
			ent = utils.FilterChar(entry, ":", false)
			if len(ent) > 1 {
				log.Println(string(ent[1]))
				if string(ent[1]) == "/" {

					bleep = append(bleep, ent[2:])
				}
			}

		}
		utils.Err(err)
	}
}
