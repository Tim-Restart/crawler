package main

import (
	"strings"
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"log"
	"net/url"
)

func GetURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {

	// get the URL's from the HTML here

	// parse the URL data to break it down into nodes
	// Nodes are a type as per below:
	/*
	type Node struct {
	Parent, FirstChild, LastChild, PrevSibling, NextSibling *Node

	Type      NodeType
	DataAtom  atom.Atom
	Data      string
	Namespace string
	Attr      []Attribute
}

	*/

	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Println("Error parsing baseURL string")
		return nil, err
	}

	links := make([]string, 10)

	htmmlReader := strings.NewReader(htmlBody)
	nodeTree, err := html.Parse(htmmlReader)
	if err != nil {
		fmt.Println("Error parsing HTML data to nodes")
		log.Fatal(err)
	}

	for n := range nodeTree.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			for _, a := range n.Attr {
				if a.Key == "href" {
					// Check if a.Val has a suffix here
					if strings.HasPrefix (a.Val, "http") {
						links = append(links, a.Val)
					} else {
						relativeURL, err := url.Parse(a.Val)
						if err != nil {
							fmt.Println("Error parsing relative URL string")
							return nil, err
						}
						finalURL := baseURL.ResolveReference(relativeURL)
						links = append(links, finalURL.String())
						break

					}
					
					
					break
				}
			}
		}
	}

	return links, nil


}