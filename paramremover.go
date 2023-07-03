package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
)

func removeParameters(u *url.URL) *url.URL {
	u.RawQuery = ""
	return u
}

func filterUniqueURLs(urls []string) []string {
	uniqueURLs := make(map[string]bool)
	var unique []string

	for _, url := range urls {
		uniqueURLs[url] = true
	}

	for url := range uniqueURLs {
		unique = append(unique, url)
	}

	return unique
}

func processURLs(urls []string) []string {
	var processedURLs []string

	for _, rawURL := range urls {
		u, err := url.Parse(rawURL)
		if err != nil {
			log.Printf("Error parsing URL %s: %v", rawURL, err)
			continue
		}

		cleanURL := removeParameters(u)
		processedURLs = append(processedURLs, cleanURL.String())
	}

	return processedURLs
}

func main() {
	urls := []string{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		url := scanner.Text()
		urls = append(urls, url)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	processedURLs := processURLs(urls)
	uniqueURLs := filterUniqueURLs(processedURLs)

	for _, url := range uniqueURLs {
		fmt.Println(url)
	}
}
