package main 

import (
	"fmt"
	"os"
	"net/url"
	"sync"
	"golang.org/x/net/html"
	)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}



func main(){

	maxConcurrency := 1

	var website string

	if len(os.Args) <= 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} else {
		fmt.Printf("starting crawl of: %v\n", os.Args[1])
		website = os.Args[1]
	}

	baseURLParsed, err := html.Parse(website)
	if err != nil {
		fmt.Println("Error parsing base URL")
		return err
	}
	
	// Pickup here with the Mu and channels stuff
	cfg := &config{
		pages: make(map[string]int),
		baseURL: baseURLParsed,
		mu: &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg : &sync.WaitGroup{},
	}

	cfg.wg.Add(1)
	go func (cfg *config) {
		defer cfg.wg.Done()
		cfg.concurrencyControl <- struct{}{}
		defer func() { <- cfg.concurrencyControl }()
		cfg.crawlPage(website) 	
	} (cfg)
	cfg.wg.Wait()

	cfg.mu.Lock()
	for _, page := range cfg.pages {
		fmt.Println(page)
	}
	cfg.mu.Unlock()
	return

	
}
