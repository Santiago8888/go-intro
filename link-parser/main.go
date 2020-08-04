package main

import (
	"os"
	"log"
	"fmt"
	"golang.org/x/net/html"
)


func main() {
	r, err := os.Open("ex1.html")

	if err != nil {
		log.Fatalln("Couldn't open the html file", err)
	}

	t := ""
	z := html.NewTokenizer(r)

	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			return
		}

		if tt == html.TextToken && t == "a" {
			fmt.Println(string(z.Text()))
		}

		name, _ := z.TagName()
		if "a" == string(name) && tt == html.StartTagToken {
			_, attr, _ := z.TagAttr()
			
			fmt.Println(string(attr))
			t = "a"
		} else { t = "" }

	}
}

