package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fmt"
	"os"
)


type Acts struct {
	acts map[string] interface {} `json:"-"`
}


func main() {
	mux := serve()

	fmt.Println("Starting Go Adventure server on http://localhost:8080")

	// Open our jsonFile
	jsonFile, err := os.Open("gopher.json")

	// if we os.Open returns an error then handle it
	if err != nil {
	    fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	jsonFile.Close()

	parseStory(byteValue, "intro")

	http.ListenAndServe(":8080", mux)
}


func serve() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/intro", hello)
	return mux
}


func hello(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "Hello, world!")
	fmt.Fprintln(w, r.RequestURI)
}



func parseStory(byteValue []byte, name string) {
	f := Acts{}

	if err := json.Unmarshal([]byte(byteValue), &f.acts); err != nil {
		panic(err)
	}


	intro := f.acts[name]
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
}

