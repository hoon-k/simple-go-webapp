package main

import (
    "net/http"
    "html/template"
)

type item struct {
    ID      int
    Name    string
    Price   float64
}

type inventory struct {
    ID      string
    Name    string
    Items   []item
}

var baseTemplates = []string {
    "templates/master.html",
    "templates/header.html",
    "templates/footer.html",
    "templates/nav.html",
}

var inventoryData = make(map[string] inventory)

func renderPage(w http.ResponseWriter, templateFile string, data interface{}) {
    files := append(baseTemplates, templateFile)
    templates := template.Must(template.ParseFiles(files...))
    templates.ExecuteTemplate(w, "master", data)
}

func registerHandlers() {
    http.HandleFunc("/", home)
    http.HandleFunc("/home", home)
    http.HandleFunc("/detail", detail)
    http.HandleFunc("/inventory", inventoryHandler)
}

func getData() {
    inventoryData["1"] = inventory {
        ID: "1",
        Name: "Fruits",
        Items: []item {
            item { 1, "Apple", 2.99 },
            item { 2, "Pear", 3.99 },
            item { 3, "Orange", 4.99 },
        },
    }

    inventoryData["2"] = inventory {
        ID: "2",
        Name: "Vegetables",
        Items: []item {
            item { 1, "Carrots", 2.99 },
            item { 2, "Cabbage", 3.99 },
            item { 3, "Cucumber", 4.99 },
        },
    }

    inventoryData["3"] = inventory {
        ID: "3",
        Name: "Grains",
        Items: []item {
            item { 1, "Oats", 2.99 },
            item { 2, "Rice", 3.99 },
            item { 3, "Wheat", 4.99 },
        },
    }
}

func main() {
    files := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", files))
    registerHandlers()
    getData()
    http.ListenAndServe(":8080", nil)
}