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

	for i, k := range v1 {
		if i == 0 {
			fmt.Printf("STORY: %+v\n", k)
		}
	}


	options := v["options"]
	v2 := options.([]interface{})

	fmt.Println("OPTIONS:")
	for i, k := range v2 {
		option := k.(map[string]interface{})
		arc := option["arc"]
		fmt.Printf("%d. %+v\n", i + 1, arc)
	}


	defer jsonFile.Close()
}

