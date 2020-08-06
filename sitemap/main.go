package main

import (
	"fmt"
	"strings"

	"./parser"
)

func main() {

	url := "https://www.calhoun.io"
	links := fetchLinks(url, url)
	// fmt.Println(links)

	visitedLinks := []string{url + "/"}
	for _, link := range links[:5] {
		if !isVisited(link.Href, visitedLinks){
			// visitedLinks = append(visitedLinks, link.Href)
			fmt.Println(link.Href)
			fmt.Println(fetchLinks(link.Href, url))
			visitedLinks = append(visitedLinks, link.Href)
		}
	}

	fmt.Println(visitedLinks)
}

func isVisited(url string, visitedLinks []string) bool {
    for _, link := range visitedLinks {
        if url == link {
            return true
        }
    }
    return false
}

func getUniqueLinks(links[]parser.Link) []parser.Link{
	// append([]int{1,2}, []int{3,4}...)
	uniqueLinks := make([]parser.Link, 0)

	for _, link := range links {
		// fmt.Println(link.Text)
		isUnique := true
		for _, uniqueLink := range uniqueLinks {
			// fmt.Println(uniqueLink.Text)
			if link.Href == uniqueLink.Href {
				isUnique = false
				break
			}
		}

		if isUnique == true {
			uniqueLinks = append(uniqueLinks, link)
		}

	}

	return uniqueLinks
}

func fetchLinks(url, home string) []parser.Link {
	links := parser.Parse(url)
	// fmt.Println(links)
	
	domainLinks := make([]parser.Link, 0)
	for _, link := range links {
		// fmt.Println(link)
		if strings.HasPrefix(string(link.Href[0]), "/"){
			// fmt.Println(link.Href)
			link.Href = home + link.Href
			domainLinks = append(domainLinks, link)
		} else if strings.HasPrefix(string(link.Href), url){
			// fmt.Println(link.Href)
			domainLinks = append(domainLinks, link)
		}
	}

	uniqueLinks := getUniqueLinks(domainLinks)
	return uniqueLinks
}

