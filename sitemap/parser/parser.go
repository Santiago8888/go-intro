package parser

import (
	"log"
	// "fmt"
	"golang.org/x/net/html"

	"io/ioutil"
	"net/http"
	"bytes"
)


type Link struct {
  Href string
  Text string
}


func Parse(url string) []Link {
	resp, e := http.Get(url)

	if e != nil {
		log.Fatalln("Couldn't fetch url", e)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	r := bytes.NewReader(body)

	link := Link{}
	links := make([]Link, 0)

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
		} else if "a" == string(name) && tt == html.EndTagToken { 
			// fmt.Println(link)
			links = append(links, link)
		} else {
			t = "" 
		}
	}

	return links
}


