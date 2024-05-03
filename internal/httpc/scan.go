package httpc

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
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

		var count float64 = float64(len(bleep))
		var done float64
		for _, entry := range bleep {
			resp, _ := http.Get(scan.Host + "/" + entry)
			done++

			fmt.Print("\033[2K\033[999D")
			if resp.StatusCode != 404 {

				fmt.Printf("[%v] /%v\n", resp.StatusCode, entry)
			}

			fmt.Printf("Progess: %v%v (%v/%v)", math.Round(done/count*1000)/10, "%", done, count)
		}
	} else {
		file, err := os.Open(scan.Wordlist)
		utils.Err(err)
		stat, err := file.Stat()
		bufSize := stat.Size()
		utils.Err(err)
		defer file.Close()

		scanner := bufio.NewScanner(file)

		var count int
		for scanner.Scan() {
			count++
		}

		var done int

		buf := make([]byte, bufSize)
		scanner.Buffer(buf, int(bufSize))
		// optionally, resize scanner's capacity for lines over 64K, see next example
		for scanner.Scan() {
			resp, _ := http.Get(scan.Host + "/" + scanner.Text())

			done++

			if resp.StatusCode != 404 {

				fmt.Printf("[%v] %v/%v\n", resp.StatusCode, scan.Host, scanner.Text())
			}

			fmt.Printf("Progess: %v (%v/%v)", math.Round(float64(done/count)), done, count)
		}
	}

	return nil
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
				}
			}

		}
		utils.Err(err)
	} else if status == 404 {
		log.Fatalln("robots.txt does not exist on this host")
	}
}
