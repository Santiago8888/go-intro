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


	intro := f.acts["intro"]
	v := intro.(map[string]interface{})
	fmt.Printf("TITLE: %+v\n", v["title"])

	story := v["story"]
	v1 := story.([]interface{})
	// fmt.Println(v1)

	for i, k := range v1 {
		if i == 0 {
			fmt.Printf("STORY: %+v\n", k)
		}
	}


	defer jsonFile.Close()
}

