package main

import (
	"fmt"
	"./parser"
)

func main() {
	links := parser.Parse("http://example.com/")
	fmt.Println(links[0].Text)
	fmt.Println(links[0].Href)
}

