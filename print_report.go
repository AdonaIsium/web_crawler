package main

import (
	"fmt"
	"sort"
)

func printReport(pages map[string]int, baseURL string) {
	fmt.Printf("REPORT for %s\n", baseURL)

	type page struct {
		pageURL string
		links   int
	}
	var sortedPages []page
	for k, v := range pages {
		sortedPages = append(sortedPages, page{k, v})
	}

	sort.Slice(sortedPages, func(i, j int) bool {
		return sortedPages[i].links > sortedPages[j].links
	})

	for _, p := range sortedPages {
		fmt.Printf("Found %d internal links to %s\n", p.links, p.pageURL)
	}
}
