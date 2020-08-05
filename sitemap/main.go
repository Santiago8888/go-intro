package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
)

func main() {
	// resp, err := http.Get("https://www.calhoun.io/")
	resp, err := http.Get("http://example.com/")

	if err != nil {
		log.Fatalln("Couldn't fetch url", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

