package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	base, err := url.Parse(rawBaseURL)
	if err != nil {
		return
	}
	curr, err := url.Parse(rawCurrentURL)
	if err != nil {
		return
	}
	if base.Hostname() != curr.Hostname() {
		return
	}

	nURL, err := normalizeURL(curr.String())
	if err != nil {
		return
	}

	if _, ok := pages[nURL]; ok {
		pages[nURL]++
		return
	}

	pages[nURL] = 1

	fmt.Printf("currently crawling: %s\n", curr.String())
	htmlBody, err := getHTML(curr.String())
	if err != nil {
		return
	}

	foundURLs, err := getURLsFromHTML(htmlBody, base.String())
	if err != nil {
		return
	}

	for _, link := range foundURLs {
		crawlPage(rawBaseURL, link, pages)
	}

}
