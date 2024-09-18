package httpc

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"sync"

	"github.com/adzsx/difugo/internal/utils"
)

var (
	scan  utils.Input
	bleep []string
	count float64
	done  float64
	wg    sync.WaitGroup
	jobs  []chan string

	// Set up scanner to not follow redirects
	scanner http.Client = http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse }}
)

func Scan(input utils.Input) error {
	scan = input

	scanner = http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	if scan.Robots {
		utils.Verbose(1, "Talking to robots...")
		Robots(scan.Host)

		count = float64(len(bleep))
		jobs = append(jobs, make(chan string, int(count)))

		for _, entry := range bleep {
			jobs[0] <- entry + scan.Suffix
			wg.Add(1)
		}

	} else {
		utils.Verbose(1, "Started scan")
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
		jobs = append(jobs, make(chan string, int(count)))

		utils.Verbose(1, "Wordlist loaded")

		for scanner.Scan() {
			jobs[0] <- scanner.Text() + scan.Suffix
			wg.Add(1)
		}

	}

	for i := 0; i < scan.Workers; i++ {
		go worker(jobs)
	}

	wg.Wait()

	return nil
}

func worker(jobs []chan string) {
	for n := range jobs[0] {
		GetPath(n)
		wg.Done()
	}
}

func GetPath(path string) {
	resp, err := scanner.Get(scan.Host + path)
	if err != nil {
		log.Println(err)
		return
	}
	done++

	var redirect string
	if resp.StatusCode > 299 && resp.StatusCode < 399 {
		// Redirect

		url, err := resp.Location()
		utils.Err(err)
		redirect = url.Path
	}

	if len(scan.StatShow) > 0 {
		if utils.InIntSl(scan.StatShow, resp.StatusCode) {
			fmt.Print("\033[2K\033[999D")
			fmt.Printf("[%v] /%v", resp.StatusCode, path)
			if redirect != "" {
				fmt.Printf("	-> (%v)", redirect)
			}
			fmt.Printf("\nProgess: %v%v (%v/%v)", math.Round(done/count*1000)/10, "%", done, count)
		}

		// Check if return code is in Ignore codes
	} else if !utils.InIntSl(scan.StatHide, resp.StatusCode) {
		fmt.Print("\033[2K\033[999D")

		fmt.Printf("[%v] /%v", resp.StatusCode, path)
		if redirect != "" {
			fmt.Printf("	-> (%v)", redirect)
		}
		fmt.Printf("\nProgess: %v%v (%v/%v)", math.Round(done/count*1000)/10, "%", done, count)
	}

	if int(done)%10 == 0 {
		fmt.Print("\033[2K\033[999D")
		fmt.Printf("Progess: %v%v (%v/%v)", math.Round(done/count*1000)/10, "%", done, count)
	}
}
