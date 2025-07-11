package main

import (
	"net/http"
	"io"
	"log"
	"fmt"
	"strings"
	)

func GetHTML(rawURL string) (string, error) {

	res, err := http.Get(rawURL)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 399 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if !strings.HasPrefix(res.Header.Get("Content-Type"), "text/html") {
		fmt.Println("Header not text/html")
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	return string(body), nil

}


func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {

	// Base and Current are the same for the first
	// Current is used to do the calls, base is used for a base case
	crawlURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println("Error normalising URL")
		os.Exit(1)
	}

	// Uses url.Parse to compare the baseURL and CurrentURL (after normalization)
	// to make sure we haven't left the page
	err = compareURL(rawBaseURL, rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Start the HTML collection and review

	html, err := GetHTML(crawlURL)
	if err != nil {
		fmt.Println("Error getting html")
		os.Exit(1)
	}

	crawlURL, exists := pages[link]
	if !exists {
		pages[crawlURL] = 1
	}

	links, err2 := GetURLsFromHTML(html, crawlURL)
	if err2 != nil {
		fmt.Println("Error getting links from HTML")
		os.Exit(1)
	}
	
	for i := range links {
		page, exists := pages[links[i]]
		if exists {
			pages[links[i]]:++
		} else {
			pages[links[i]] = 1
		}
	}

	
	


}