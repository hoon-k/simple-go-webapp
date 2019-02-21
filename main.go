package main

import (
    "fmt"
    "net/http"
    "html/template"
)

// Item struct
type Item struct {
    Name    string
    Price   float64
}

func headers(w http.ResponseWriter, r *http.Request) {
    h := r.Header
    fmt.Fprintln(w, h)
}

func home(w http.ResponseWriter, r *http.Request) {
    someData := []Item {
        Item { "Apple", 2.99 },
        Item { "Pear", 3.99 },
        Item { "Orange", 4.99 },
    }

    templates := template.Must(template.ParseGlob("templates/*.html"))
    templates.ExecuteTemplate(w, "master", someData)
}

func main() {
     files := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", files))

    http.HandleFunc("/home", home)
    http.ListenAndServe(":8080", nil)
}