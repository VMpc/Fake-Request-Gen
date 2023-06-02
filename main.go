package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os/exec"
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
		urlUsage string = "(MUST return a raw csv file) Sets the specified URL"

		breakValue int    = 60
		breakUsage string = "Set the maximum amount of time the program will view pages for in seconds"

		viewValue int    = 60
		viewUsage string = "Set the maximum amount of time it views a page for in seconds"
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
		fmt.Println("Going to ", url)

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

/* Scrapes top 500 sites from https://moz.com */
func ScrapeData() (Urls []string, err error) {
	res, err := http.Get(URL)
	if err != nil {
		return
	}

	data, err := csv.NewReader(res.Body).ReadAll()
	if err != nil {
		return
	}

	for _, line := range data {
		Urls = append(Urls, "https://"+line[1])
		fmt.Println(line[1])
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
)
