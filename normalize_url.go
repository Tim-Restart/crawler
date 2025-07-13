package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(link string) (string, error) {

	// url input is the string that needs to be sanatised
	// An example of a normalized url is : blog.boot.dev/path
	// Inital thoughts are to detect and remove prefixes for http/https
	// Suffix to remove any trailing /

	normalUrl, err := url.Parse(link)
	if err != nil {
		fmt.Println("Error parsing URL string")
		return "", err
	}

	sanatised := normalUrl.Host + normalUrl.Path
	fmt.Println(sanatised)
	return strings.TrimSuffix(sanatised, "/"), nil
}

// Helper function for comparing URLS in the crawlPage function

func compareURL(baseURL *url.URL, currentURL string) error {

	crawlURLCheck, err := url.Parse(currentURL)
	if err != nil {
		return err
	}

	if baseURL.Host != crawlURLCheck.Host {
		err = fmt.Errorf("Hosts do not match, left crawl target page")
		return err
	} else {
		return nil
	}
}

// func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool)

func stringToURL(link string) (*url.URL, error) {
	baseURLParsed, err := url.Parse(link)
	if err != nil {
		fmt.Println("Error parsing base URL")
		return nil, err
	}
	return baseURLParsed, nil
}
