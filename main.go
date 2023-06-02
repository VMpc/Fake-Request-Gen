package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

func init() {
	const (
		browserValue string = "firefox"
		browserUsage string = "Sets the browser to use"

		browserArgsValue string = "--headless --private-window"
		browserArgsUsage string = "Sets the browser args to use"

		urlValue string = "https://moz.com/top-500/download/?table=top500Domains"
		urlUsage string = "Sets the specified URL/File path"

		breakValue int    = 60
		breakUsage string = "Sets the maximum amount of time the program will view pages for in seconds"

		viewValue int    = 60
		viewUsage string = "Sets the maximum amount of time it views a page for in seconds"
	)

	flag.StringVar(&Browser, "browser", browserValue, browserUsage)
	flag.StringVar(&BrowserArgs, "browserargs", browserArgsValue, browserArgsUsage)

	flag.StringVar(&URL, "url", urlValue, urlUsage)

	flag.IntVar(&breakMax, "breaktime", breakValue, breakUsage)
	flag.IntVar(&viewMax, "viewtime", viewValue, viewUsage)

	flag.StringVar(&Browser, "b", browserValue, browserUsage)
	flag.StringVar(&BrowserArgs, "ba", browserArgsValue, browserArgsUsage)

	flag.StringVar(&URL, "u", urlValue, urlUsage)

	flag.IntVar(&breakMax, "bt", breakValue, breakUsage)
	flag.IntVar(&viewMax, "vt", viewValue, viewUsage)

	flag.Parse()
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	urls, err := ScrapeData()
	if err != nil {
		fmt.Printf("Could not scrap data: %s\n", err)
		return
	}
	fmt.Printf("Found %d urls, starting program\n", len(urls))

	for {
		url := urls[rand.Intn(len(urls))]
		args := append(strings.Fields(BrowserArgs), url)

		cmd := exec.Command(Browser, args...)
		if err := cmd.Start(); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Going to", url)

		/* View the page for X amount of seconds then kill firefox */
		time.Sleep(time.Duration(rand.Intn(viewMax)) * time.Second)
		if err := cmd.Process.Kill(); err != nil {
			fmt.Println(err)
		}

		/* "Random" choice between a short sleep and a long sleep */
		if rand.Intn(100) >= 50 {
			time.Sleep(time.Duration(rand.Intn(breakMax)))
		} else {
			time.Sleep(time.Duration(rand.Intn(breakMax * 7)))
		}
	}
}

/* Handle both urls, file paths and get the URLs from it */
func parseData() ([]string, error) {
	var (
		tempData []byte
		err      error
	)

	/* A switch statement looks SLIGHTLY nicer than if/else in this situation */
	switch strings.HasPrefix(URL, "http") {
	case true:
		res, err := http.Get(URL)
		if err != nil {
			return nil, err
		}
		tempData, err = ioutil.ReadAll(res.Body)
	default:
		if _, err = os.Stat(URL); err != nil {
			return nil, err
		}
		tempData, err = os.ReadFile(URL)
	}

	return urlRegex.FindAllString(string(tempData), -1), nil
}

/* Scrapes top 500 sites from https://moz.com */
func ScrapeData() (Urls []string, err error) {
	data, err := parseData()
	if err != nil {
		return
	}

	for _, line := range data {
		if !strings.HasPrefix(line, "http") {
			line = "https://" + line
		}

		Urls = append(Urls, line)
		fmt.Println(line)
	}

	if len(Urls) == 0 {
		return Urls, fmt.Errorf("0 urls found")
	}

	return
}

var (
	Browser     string
	BrowserArgs string
	URL         string
	breakMax    int
	viewMax     int
	// Gladly stolen from https://regex101.com/library/hM4wG0
	urlRegex *regexp.Regexp = regexp.MustCompile(`(?:https?:\/\/)?(?:[-\w_\.]{2,}\.)?([-\w_]{1,}\.[a-z]{2,4})(?:\/\S*)?`)
)
