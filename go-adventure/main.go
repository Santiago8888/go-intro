package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"os"
)


type Acts struct {
	acts map[string] interface {} `json:"-"`
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


	// fmt.Printf("%+v\n", f.acts["intro"])

	intro := f.acts["intro"]
	v := intro.(map[string]interface{})
	fmt.Printf("%+v\n", v["title"])


	defer jsonFile.Close()
}


/*
	keys := make([]string, 0, len(f.acts))
	for k := range f.acts {
		keys = append(keys, k)
	}
	fmt.Println(keys)
	fmt.Println("\nSuccessfully Opened gopher.json")

*/


