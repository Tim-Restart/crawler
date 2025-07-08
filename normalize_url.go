package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(link string) (string, error){

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