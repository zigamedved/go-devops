package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func getUrlFromArgs(args []string) (string, error) {
	url, err := url.ParseRequestURI(args[1])
	if err != nil {
		fmt.Printf("Error while parsing URI!")
		return "", err
	}

	return url.String(), nil
}

func main() {
	args := os.Args

	if len(args) <= 1 {
		fmt.Printf("Usage: go run main.go <url>\n")
		os.Exit(1)
	}

	url, err := getUrlFromArgs(os.Args)
	if err != nil {
		fmt.Printf("Couldn't get url from args!")
		os.Exit(1)
	}

	res, err2 := http.Get(url)
	if err2 != nil {
		log.Fatal(err2)
	}
	defer res.Body.Close()

	body, err3 := io.ReadAll(res.Body)
	if err3 != nil {
		log.Fatal(err3)
	}

	fmt.Printf("HTTP Status Code %d\nBody: %s\n", res.StatusCode, body)
}
