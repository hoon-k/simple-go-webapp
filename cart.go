package main

import (
    "net/http"
    // "strconv"
)

type cartItem struct {
    ID          string
    unitItem    item
    quantity    int
}

var cart = make(map[string] map[string] cartItem)

func cartHandler(w http.ResponseWriter, r *http.Request) {
    renderPage(w, "templates/cart.html", cart)
}

// func addItem(w http.ResponseWriter, r *http.Request) {
//     id := r.URL.Query().Get("id")
//     q, _ := strconv.ParseInt(r.URL.Query().Get("q"), 0, 64)
//     renderPage(w, "templates/cart.html", cart)
// }
