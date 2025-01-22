package main

import (
	"ChapterFive/links"
	// "encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func createDoc(item string) {
	resp, err := http.Get(item)
	if err != nil {
		fmt.Printf("Error making request: %s", err)
	}
	if resp.StatusCode != 200 {
		resp.Body.Close()
		fmt.Printf("Request failed with status %s", resp.Status)
	}

	// var body io.ReadCloser
	// if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
	// 	resp.Body.Close()
	// 	fmt.Printf("Error getting response: %s", err)
	// }
	defer resp.Body.Close()

	file, err := os.Create(item)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating file %s: %v\n", item, err)
	}

	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", item, err)
		os.Exit(1)
	}
}

func sanitizeUrl(url url.URL) string {
	sanitized := url.Path
	if sanitized == "" || sanitized == "/" {
		return "index.html"
	}
	sanitized = strings.ReplaceAll(sanitized, "/", "_")
	sanitized = strings.ReplaceAll(sanitized, "?", "_")
	sanitized = strings.ReplaceAll(sanitized, ":", "_")
	return sanitized
}

func crawl(item string) []string {
	fmt.Print(item)
	list, err := links.Extract(item)
	if err != nil {
		log.Print(err)
	}
	baseUrl, err := url.Parse(item)
	if err != nil {
		log.Print(err)
	}
	fileName := sanitizeUrl(*baseUrl)
	host := baseUrl.Host
	for _, item := range list {
		parsedURL, err := url.Parse(item)
		if err != nil {
			log.Fatal(err)
		}

		if host == parsedURL.Host {
			createDoc(fileName)
		}

	}
	return list
}
