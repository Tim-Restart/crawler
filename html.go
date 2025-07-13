package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func GetHTML(rawURL string) (string, error) {

	res, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 399 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if !strings.HasPrefix(res.Header.Get("Content-Type"), "text/html") {
		err = fmt.Errorf("Header not text/html")
		return "", err
	}
	if err != nil {
		return "", err
	}
	return string(body), nil

}

func (cfg *config) crawlPage(rawCurrentURL string) {

	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	// Base and Current are the same for the first
	// Current is used to do the calls, base is used for a base case
	crawlURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println("Error normalising URL")
		os.Exit(1)
	}

	// Uses url.Parse to compare the baseURL and CurrentURL (after normalization)
	// to make sure we haven't left the page
	err = compareURL(cfg.baseURL, rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Start the HTML collection and review
	if !cfg.addPageVisit(crawlURL) {
		return
	}

	html, err := GetHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error getting html for %v\n", rawCurrentURL)

	}

	links, err2 := GetURLsFromHTML(html, rawCurrentURL)
	if err2 != nil {
		fmt.Println("Error getting links from HTML")
		os.Exit(1)
	}

	for _, newLink := range links {
		cfg.wg.Add(1)
		go cfg.crawlPage(newLink)
	}

}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, exists := cfg.pages[normalizedURL]; exists {
		cfg.pages[normalizedURL]++
		return false
	}

	cfg.pages[normalizedURL] = 1
	return true
}
