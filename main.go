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
		fmt.Printf("starting crawl of: %s\n", args[1])
		pages := make(map[string]int)
		crawlPage(args[1], args[1], pages)
	}
}
