package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

// {"page":"words","input":"","words":[]}
type words struct {
	Page  string   `json:"page"`
	Input string   `json:"input"`
	Words []string `json:"words"`
}

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
	if res.StatusCode != 200 {
		fmt.Printf("Response status code is not 200 but %d\n", res.StatusCode)
		os.Exit(1)
	}

	body, err3 := io.ReadAll(res.Body)
	if err3 != nil {
		log.Fatal(err3)
	}

	var words words
	err4 := json.Unmarshal(body, &words)
	if err4 != nil {
		fmt.Printf("Error while parsing JSON")
		os.Exit(1)
	}

	fmt.Printf("Parsed JSON\nPage: %s\nInput: %s, \nWords: %v\n", words.Page, words.Input, words.Words)
}
