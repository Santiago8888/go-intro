package main

import (
	"encoding/json"
//	"html/template"
	"io/ioutil"
	"net/http"
	"fmt"
	"os"
)


type Acts struct {
	acts map[string] interface {} `json:"-"`
}


var f = Acts{}

func main() {


	// Open our jsonFile
	jsonFile, err := os.Open("gopher.json")

	// if we os.Open returns an error then handle it
	if err != nil {
	    fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	jsonFile.Close()


	if err := json.Unmarshal([]byte(byteValue), &f.acts); err != nil {
		panic(err)
	}


	mux := serve()
	fmt.Println("Starting Go Adventure server on http://localhost:8080/intro")
	http.ListenAndServe(":8080", mux)
}


func serve() *http.ServeMux {
	mux := http.NewServeMux()

	// TODO: Add default route that maps to index.
	// mux.HandleFunc("/", route)

	for i := range f.acts {
		mux.HandleFunc("/"+ i, route)
	}

	return mux
}


func trimFirstChar(s string) string {
    for i := range s {
        if i > 0 {
            return s[i:]
        }
    }

    return s[:0]
}

func route(w http.ResponseWriter, r *http.Request) {
	title := trimFirstChar(r.RequestURI)
	parseStory(title, w)
}



func parseStory(name string, w http.ResponseWriter) {
	intro := f.acts[name]
	v := intro.(map[string]interface{})

	title := v["title"]
	fmt.Fprintln(w, "<h1>" + title.(string) + "</h1>")
	fmt.Fprintln(w, "\n")


	story := v["story"]
	v1 := story.([]interface{})

	for _, paragraph := range v1 {
		fmt.Fprintln(w, "<p>" + paragraph.(string) + "</p>")
	}

	fmt.Fprintln(w, "\n")

	options := v["options"]
	v2 := options.([]interface{})
	for _, k := range v2 {
		option := k.(map[string]interface{})
		arc := option["arc"]
		fmt.Fprintln(w, `<a href="` + arc.(string) + `">` + arc.(string) + "</a><br/>")
	}
}

