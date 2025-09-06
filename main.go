package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("no website provided\n")
		os.Exit(1)
	} else if len(args) > 2 {
		fmt.Printf("too many arguments provided\n")
		os.Exit(1)
	} else {
		rawBaseURL := os.Args[1]

		const maxConcurrency = 3
		cfg, err := configure(rawBaseURL, maxConcurrency)
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
	}
}
