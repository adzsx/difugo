package httpc

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/adzsx/dirsgover/internal/utils"
)

var (
	scan    utils.Input
	bleep   []string
	count   float64
	done    float64
	wg      sync.WaitGroup
	jobs    chan string
	status  int
	workers int
	client  http.Client = http.Client{Timeout: 3 * time.Second}
)

func Scan(input utils.Input) error {
	scan = input

	if scan.Robots {
		Robots(scan.Host)

		count = float64(len(bleep))
		jobs = make(chan string, int(count))

		for i := 0; i < scan.Workers; i++ {
			go worker(jobs)
			workers++
		}

		for _, entry := range bleep {
			jobs <- entry
		}

		close(jobs)

	} else {
		file, err := os.Open(scan.Wordlist)
		utils.Err(err)
		stat, err := file.Stat()
		bufSize := stat.Size()
		utils.Err(err)
		defer file.Close()

		scanner := bufio.NewScanner(file)

		buf := make([]byte, bufSize)
		scanner.Buffer(buf, int(bufSize))

		count = float64(utils.LineCount(scan.Wordlist))
		jobs = make(chan string, int(count))

		for i := 0; i < scan.Workers; i++ {
			go worker(jobs)
			workers++
		}

		for scanner.Scan() {
			wg.Add(1)
			jobs <- scanner.Text()
		}

	}

	wg.Wait()

	return nil
}

func worker(jobs <-chan string) {
	for n := range jobs {
		if n[0:1] != "~" {
			GetPath(n)
			wg.Done()
		}
	}
}

func GetPath(path string) {
	resp, err := client.Get(scan.Host + "/" + path)
	if err != nil {
		return
	}
	done++

	if resp.StatusCode == 429 {
		log.Println("Err 429")
	}

	if len(scan.StatShow) > 0 {
		if utils.InIntSl(scan.StatShow, resp.StatusCode) {
			fmt.Printf("\033[2K\033[999D[%v] /%v\nProgess: %v%v (%v/%v)", resp.StatusCode, path, math.Round(done/count*1000)/10, "%", done, count)
		}
	} else if !utils.InIntSl(scan.StatHide, resp.StatusCode) {
		fmt.Printf("\033[2K\033[999D[%v] /%v\nProgess: %v%v (%v/%v)", resp.StatusCode, path, math.Round(done/count*1000)/10, "%", done, count)
	}

	if int(done)%10 != -1 {
		fmt.Print("\033[2K\033[999D")
		fmt.Printf("Progess: %v%v (%v/%v)", math.Round(done/count*1000)/10, "%", done, count)
	}
}
