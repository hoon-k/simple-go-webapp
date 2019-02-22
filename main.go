package main

import (
    "net/http"
    "html/template"
    "strconv"
)

type item struct {
    ID      int
    Name    string
    Price   float64
}

var baseTemplates = []string {
    "templates/master.html",
    "templates/header.html",
    "templates/footer.html",
    "templates/nav.html",
}

var someData = []item {
    item { 1, "Apple", 2.99 },
    item { 2, "Pear", 3.99 },
    item { 3, "Orange", 4.99 },
}

func home(w http.ResponseWriter, r *http.Request) {
    renderPage(w, "templates/home.html", someData)
}

func detail(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.ParseInt(r.URL.Query().Get("id"), 0, 64)
    renderPage(w, "templates/detail.html", someData[id - 1])
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