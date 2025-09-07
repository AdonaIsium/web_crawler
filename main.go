package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("no website provided\n")
		os.Exit(1)
	} else if len(args) > 4 {
		fmt.Printf("too many arguments provided\n")
		os.Exit(1)
	} else {
		rawBaseURL := os.Args[1]

		maxConcurrency := 3
		maxPages := 10

		if len(args) >= 3 {
			concurrencyArg, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Printf("max concurrency must be int")
				os.Exit(1)
			}
			maxConcurrency = concurrencyArg
		}

		if len(args) == 4 {
			pagesArg, err := strconv.Atoi(args[3])
			if err != nil {
				fmt.Printf("max pages must be int")
				os.Exit(1)
			}
			maxPages = pagesArg
		}

		cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
		if err != nil {
			fmt.Printf("error configure: %v", err)
			os.Exit(1)
		}

		cfg.wg.Add(1)
		go cfg.crawlPage(rawBaseURL)
		cfg.wg.Wait()

		for normalizedURL, count := range cfg.pages {
			fmt.Printf("%d - %s\n", count, normalizedURL)
		}

		printReport(cfg.pages, cfg.baseURL.String())
	}
}
