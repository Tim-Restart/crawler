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

	body, err :=GetHTML(website)
	if err != nil {
		fmt.Println("Error parsing website body")
		os.Exit(1)
	}

	fmt.Println(body)

	
}
