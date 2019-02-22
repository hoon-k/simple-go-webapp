package main

import (
    "net/http"
    "html/template"
)

// Item struct
type Item struct {
    Name    string
    Price   float64
}

var baseTemplates = []string {
    "templates/master.html",
    "templates/header.html",
    "templates/footer.html",
    "templates/nav.html",
}

var someData = []Item {
    Item { "Apple", 2.99 },
    Item { "Pear", 3.99 },
    Item { "Orange", 4.99 },
}

func home(w http.ResponseWriter, r *http.Request) {
    renderPage(w, "templates/home.html", someData)
}

func detail(w http.ResponseWriter, r *http.Request) {
    renderPage(w, "templates/detail.html", someData[0])
}

func renderPage(w http.ResponseWriter, templateFile string, data interface{}) {
    files := append(baseTemplates, templateFile)
    templates := template.Must(template.ParseFiles(files...))
    templates.ExecuteTemplate(w, "master", data)
}

func main() {
    files := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", files))

    http.HandleFunc("/", home)
    http.HandleFunc("/home", home)
    http.HandleFunc("/detail", detail)
    http.ListenAndServe(":8080", nil)
}