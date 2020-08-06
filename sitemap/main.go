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

	uniqueLinks := make([]parser.Link, 0)
	for _, link := range links {
		uniqueLinks = append(uniqueLinks, link)
	}

	visitedLinks := []string{url + "/"}
	i := 0
	for {
		link := uniqueLinks[i]
		fmt.Println(len(uniqueLinks))
		if !isVisited(link.Href, visitedLinks){
			// visitedLinks = append(visitedLinks, link.Href)

			// fmt.Println("Link to Visit: ")
			// fmt.Println(link.Href)
			newLinks := fetchLinks(link.Href, url)
			// fmt.Println("Fetched Links: ")
			// fmt.Println(newLinks)

			// fmt.Println("Fetched Links Length: ")
			// fmt.Println(len(newLinks))

			// fmt.Println("Len of Links I: ")
			// fmt.Println(len(uniqueLinks))
			uniqueLinks = getUniqueLinks(append(uniqueLinks, newLinks...))

			// fmt.Println("Updated Links: ")
			// fmt.Println(uniqueLinks)

			// fmt.Println("Len of Links II: ")
			// fmt.Println(len(uniqueLinks))

			visitedLinks = append(visitedLinks, link.Href)
			// fmt.Println("Visited Links: ")
			// fmt.Println(visitedLinks)

			links := make([]parser.Link, 0)
			for _, link := range uniqueLinks {
				links = append(links, link)
			}
		}

		i ++
		if i + 1 == len(uniqueLinks){
			break
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

