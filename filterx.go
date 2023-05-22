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
	flag.Parse()

	// Get the name of the tool
	toolName := filepath.Base(os.Args[0])

	// Print the tool banner
	printBanner(toolName)

	// Get the filter name from the first argument
	if len(flag.Args()) != 1 {
		log.Fatalf("Filter name is required. Usage: cat urls.txt | %s <filter-name>", toolName)
	}
	filterName := flag.Arg(0)

	// Read the filter file
	filterFilePath := filepath.Join("/filterx", filterName+".json")
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

func printBanner(toolName string) {
	banner := `
 ________  __  __    __                         __    __ 
|        \|  \|  \  |  \                       |  \  |  \
| $$$$$$$$ \$$| $$ _| $$_     ______    ______ | $$  | $$
| $$__    |  \| $$|   $$ \   /      \  /      \ \$$\/  $$
| $$  \   | $$| $$ \$$$$$$  |  $$$$$$\|  $$$$$$\ >$$  $$ 
| $$$$$   | $$| $$  | $$ __ | $$    $$| $$   \$$/  $$$$\ 
| $$      | $$| $$  | $$|  \| $$$$$$$$| $$     |  $$ \$$\
| $$      | $$| $$   \$$  $$ \$$     \| $$     | $$  | $$
 \$$       \$$ \$$    \$$$$   \$$$$$$$ \$$      \$$   \$$
                                                         
                                                         
                                                         
	`
	fmt.Println(banner)
	fmt.Printf("Welcome to %s!\n\n", toolName)
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
