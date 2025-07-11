package main 

import (
	"fmt"
	"os"
	)



func main(){

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

	

	pages := make(map[string]int)

	crawlPage(website, website, pages)

	for _, page := range pages {
		fmt.Println(page)
	}

	

	// Testing absolute and relative URLS collections
	/*
	body, err := GetHTML(website)
	if err != nil {
		return
	}

	urls, err := GetURLsFromHTML(body, website)
	if err != nil {
		return
	}

	for i := range urls {
		fmt.Println(urls[i])
	}
	*/

	return

	
}
