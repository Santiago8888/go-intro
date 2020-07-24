package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"os"
)

type Acts struct {
	acts map[string]interface{} `json:"-"`
}

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("gopher.json")

	// if we os.Open returns an error then handle it
	if err != nil {
	    fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	f := Acts{}

	if err := json.Unmarshal([]byte(byteValue), &f.acts); err != nil {
		panic(err)
	}

	fmt.Printf("%+v", f)
	fmt.Println("\nSuccessfully Opened gopher.json")

	defer jsonFile.Close()
}

