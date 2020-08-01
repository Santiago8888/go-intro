package main

import (
	"html/template"

	"encoding/json"

	"io/ioutil"
	"net/http"
	"fmt"
	"os"
)


type Acts struct {
	acts map[string] interface {} `json:"-"`
}


type Paragraph struct {
	Text string
}

type Option struct {
	Text string
	Link string
}


type Story struct {
	Title string
	Paragraphs	[]Paragraph
	Options		[]Option
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

	tmpl := template.Must(template.ParseFiles("story.html"))

	for i := range f.acts {
		http.HandleFunc("/" + i, func(w http.ResponseWriter, r *http.Request) {
			title := trimFirstChar(r.RequestURI)
			story := createStory(title)
			tmpl.Execute(w, story)
		})
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		story := createStory("intro")
		tmpl.Execute(w, story)
	})

	fmt.Println("Starting Go Adventure server on http://localhost:8080/intro")
	http.ListenAndServe(":8080", nil)
}


func trimFirstChar(s string) string {
	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}

	return s[:0]
}


func createStory(name string) Story {
	s := Story{}

	act := f.acts[name]
	v := act.(map[string]interface{})
	title := v["title"]
	s.Title = title.(string)

	story := v["story"]
	v1 := story.([]interface{})
	paragraphs := make([]Paragraph, len(v1))
	for i, paragraph := range v1 {
		p := Paragraph{Text: paragraph.(string)}
		paragraphs[i] = p
	}
	s.Paragraphs = paragraphs

	options := v["options"]
	v2 := options.([]interface{})
	opts := make([]Option, len(v2))
	for i, k := range v2 {
		option := k.(map[string]interface{})
		arc := option["arc"]
		text := option["text"]
		o := Option{Link: arc.(string), Text: text.(string)}
		opts[i] = o
	}
	s.Options = opts

	return s
}


