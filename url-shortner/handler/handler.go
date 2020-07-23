package handler

import (
	"fmt"
	"reflect"
	"net/http"
	yaml "gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, mux http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, ok := pathsToUrls[r.RequestURI]; ok {
			http.Redirect(w, r, pathsToUrls[r.RequestURI], http.StatusSeeOther)
		}

    		mux.ServeHTTP(w, r)
  	})
}


type Keys struct {
	Path string 
	Url string
}

func getField(v *Keys, field string) string {
    r := reflect.ValueOf(v)
    f := reflect.Indirect(r).FieldByName(field)
    return string(f.String())
}

func YAMLHandler(yml []byte, mux http.Handler) http.HandlerFunc {
	keys := make([]Keys, 0)
	yaml.Unmarshal(yml, &keys)
	
	m := make(map[string]string)
	for i := range keys {
		m[getField(&keys[i], "Path")] = getField(&keys[0], "Url")
	}

	fmt.Println("map:", m)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, ok := m[r.RequestURI]; ok {
			http.Redirect(w, r, m[r.RequestURI], http.StatusSeeOther)
		}

    		mux.ServeHTTP(w, r)
  	})
}

