package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type FilterFile struct {
	Words []string `json:"words"`
}

func main() {
	// Parse command line arguments
	filterName := flag.String("g", "", "Filter name")
	filterDir := flag.String("d", "", "Directory containing filter files")
	flag.Parse()

	if *filterName == "" {
		log.Fatal("Filter name is required. Usage: filterx -g <filter-name> -d <filter-directory>")
	}

	if *filterDir == "" {
		log.Fatal("Filter directory is required. Usage: filterx -g <filter-name> -d <filter-directory>")
	}

	// Read the filter file
	filterFilePath := filepath.Join(*filterDir, *filterName+".json")
	filterData, err := ioutil.ReadFile(filterFilePath)
	if err != nil {
		log.Fatalf("Failed to read filter file: %v", err)
	}

	// Parse the filter file
	var filter FilterFile
	err = json.Unmarshal(filterData, &filter)
	if err != nil {
		log.Fatalf("Failed to parse filter file: %v", err)
	}

	// Read URLs from stdin
	urls := readUrlsFromStdin()

	// Filter and print matching URLs
	filterAndPrintUrls(urls, filter.Words)
}

func readUrlsFromStdin() []string {
	var urls []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		url := scanner.Text()
		urls = append(urls, url)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed to read URLs from stdin: %v", err)
	}

	return urls
}

func filterAndPrintUrls(urls []string, words []string) {
	for _, url := range urls {
		for _, word := range words {
			if strings.Contains(url, word) {
				fmt.Println(url)
				break
			}
		}
	}
}
