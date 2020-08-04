package main

import (
	"os"
	"log"
	"fmt"
	"golang.org/x/net/html"
)


type Link struct {
  Href string
  Text string
}


func main() {
	link := Link{}

	r, err := os.Open("ex1.html")

	if err != nil {
		log.Fatalln("Couldn't open the html file", err)
	}

	t := ""
	z := html.NewTokenizer(r)


	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}

		if tt == html.TextToken && t == "a" {
			link.Text = string(z.Text())
		}

		name, _ := z.TagName()
		if "a" == string(name) && tt == html.StartTagToken {
			_, attr, _ := z.TagAttr()
			link.Href = string(attr)
			t = "a"
		} else { t = "" }

	}

	fmt.Println(link.Text)
	fmt.Println(link.Href)
	fmt.Println(link)
}

