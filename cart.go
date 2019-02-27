package main

import (
    "net/http"
    "fmt"
    "strconv"
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

func addItem(w http.ResponseWriter, r *http.Request) {
    itemID := r.FormValue("id")
    quantity, _ := strconv.ParseFloat(r.FormValue(itemID), 64)
    unitPrice, _ := strconv.ParseFloat(r.FormValue("price"), 64)
    totalPrice := unitPrice * quantity
    fmt.Println(itemID, quantity, unitPrice, totalPrice)
    renderPage(w, "templates/cart.html", cart)
}
