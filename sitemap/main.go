package main

import (
	"fmt"
	"./parser"
)

func main() {
	links := parser.Parse("https://www.calhoun.io/")
	// fmt.Println(links[0].Text)
	// fmt.Println(links[0].Href)

	fmt.Println(links)
}

