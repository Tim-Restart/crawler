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
