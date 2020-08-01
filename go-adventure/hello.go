package main


import (
    "html/template"
    "net/http"
)



type Todo struct {
    Title string
    Done  bool
}



type TodoPageData struct {
    PageTitle string
    Todos     []Todo
}



func main() {
    tmpl := template.Must(template.ParseFiles("hello.html"))
    http.HandleFunc("/a", func(w http.ResponseWriter, r *http.Request) {
        data := TodoPageData{
            PageTitle: "My TODO list",
            Todos: []Todo{
                {Title: "Task 1", Done: false},
                {Title: "Task 2", Done: true},
                {Title: "Task 3", Done: true},
            },
        }
        tmpl.Execute(w, data)
    })

    http.HandleFunc("/b", func(w http.ResponseWriter, r *http.Request) {
        data := TodoPageData{
            PageTitle: "My TODO list",
            Todos: []Todo{
                {Title: "Task A", Done: false},
                {Title: "Task B", Done: true},
                {Title: "Task C", Done: true},
            },
        }
        tmpl.Execute(w, data)
    })

    http.ListenAndServe(":8080", nil)
}

