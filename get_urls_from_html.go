package main

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody string, base *url.URL) ([]string, error) {

	node, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	if docBase := findBaseHref(node); docBase != "" {
		if b2, err := url.Parse(docBase); err == nil {
			if b2.IsAbs() && (b2.Scheme == "http" || b2.Scheme == "https") {
				base = b2
			}
		}
	}

	var out []string
	walk(node, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					href := strings.TrimSpace(a.Val)
					if href == "" {
						continue
					}

					if isSkippableScheme(href) {
						continue
					}

					u, err := url.Parse(href)
					if err != nil {
						continue
					}
					abs := base.ResolveReference(u)

					if abs.IsAbs() && (abs.Scheme == "http" || abs.Scheme == "https") {
						out = append(out, abs.String())
					}
				}
			}
		}
	})

	return out, nil
}

func walk(n *html.Node, f func(*html.Node)) {
	if n == nil {
		return
	}
	f(n)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		walk(c, f)
	}
}

func findBaseHref(root *html.Node) string {
	var href string
	walk(root, func(n *html.Node) {
		if href != "" {
			return
		}
		if n.Type == html.ElementNode && n.Data == "base" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					href = strings.TrimSpace(a.Val)
					return
				}
			}
		}
	})
	return href
}

func isSkippableScheme(s string) bool {
	colon := strings.IndexByte(s, ':')
	if colon > 0 {
		scheme := strings.ToLower(s[:colon])
		switch scheme {
		case "http", "https":
			return false
		default:
			return true
		}
	}

	return false
}
