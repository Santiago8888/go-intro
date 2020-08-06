package main

import (
	"fmt"
	"strings"

	"./parser"
)

func main() {
	url := "https://www.calhoun.io/"
	links := parser.Parse(url)
	// fmt.Println(links)
	
	domainLinks := make([]parser.Link, 0)
	for _, link := range links {
		// fmt.Println(link)
		if strings.HasPrefix(string(link.Href[0]), "/"){
			// fmt.Println(link.Href)
			domainLinks = append(domainLinks, link)
		} else if strings.HasPrefix(string(link.Href), url){
			// fmt.Println(link.Href)
			domainLinks = append(domainLinks, link)
		}
	}

	fmt.Println(domainLinks)
}

