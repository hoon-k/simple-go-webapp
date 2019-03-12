package main

import (
    "net/http"
    "html/template"
)

type item struct {
    ID      string
    Name    string
    Price   float64
}

type inventory struct {
    ID      string
    Name    string
    Items   map[string]item
}

var baseTemplates = []string {
    "templates/master.html",
    "templates/header.html",
    "templates/footer.html",
    "templates/nav.html",
}

var inventoryData = make(map[string]inventory)
var allItems = make(map[string]item)

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
    http.HandleFunc("/cart", cartHandler)
    http.HandleFunc("/addItem", addItem)
    http.HandleFunc("/inventory-data", inventoryRawData)
}

func getData() {
    allItems["f1"] = item { "f1", "Apple", 2.99 }
    allItems["f2"] = item { "f2", "Pear", 3.99 }
    allItems["f3"] = item { "f3", "Orange", 4.99 }
    allItems["v1"] = item { "v1", "Carrots", 2.99 }
    allItems["v2"] = item { "v2", "Cabbage", 3.99 }
    allItems["v3"] = item { "v3", "Cucumber", 4.99 }
    allItems["g1"] = item { "g1", "Oats", 2.99 }
    allItems["g2"] = item { "g2", "Rice", 3.99 }
    allItems["g3"] = item { "g3", "Wheat", 4.99 }

    fruits := make(map[string]item)
    fruits["f1"] = item { "f1", "Apple", 2.99 }
    fruits["f2"] = item { "f2", "Pear", 3.99 }
    fruits["f3"] = item { "f3", "Orange", 4.99 }
    inventoryData["1"] = inventory {
        ID: "1",
        Name: "Fruits",
        Items: fruits,
    }

    veggies := make(map[string]item)
    veggies["v1"] = item { "v1", "Carrots", 2.99 }
    veggies["v2"] = item { "v2", "Cabbage", 3.99 }
    veggies["v3"] = item { "v3", "Cucumber", 4.99 }
    inventoryData["2"] = inventory {
        ID: "2",
        Name: "Vegetables",
        Items: veggies,
    }

    grains := make(map[string]item)
    grains["g1"] = item { "g1", "Oats", 2.99 }
    grains["g2"] = item { "g2", "Rice", 3.99 }
    grains["g3"] = item { "g3", "Wheat", 4.99 }
    inventoryData["3"] = inventory {
        ID: "3",
        Name: "Grains",
        Items: grains,
    }
}

func main() {
    files := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", files))
    registerHandlers()
    getData()
    http.ListenAndServe(":8080", nil)
}