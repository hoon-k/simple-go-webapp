package main

import (
    "net/http"
    "strconv"
    "math"
    "encoding/json"
)

type cartItem struct {
    ID              string
    UnitItem        item
    Quantity        int
    Total           float64
}

var cart = make(map[string]*cartItem)

func cartHandler(w http.ResponseWriter, r *http.Request) {
    renderPage(w, "templates/cart.html", cart)
}

func addItem(w http.ResponseWriter, r *http.Request) {
    itemID := r.FormValue("id")
    quantity, _ := strconv.ParseFloat(r.FormValue(itemID), 64)
    unitPrice, _ := strconv.ParseFloat(r.FormValue("price"), 64)
    totalPrice := math.Ceil((unitPrice * quantity) * 100) / 100

    if cart[itemID] == nil {
        cart[itemID] = &cartItem {
            ID: itemID,
            UnitItem: allItems[itemID],
            Quantity: int(quantity),
            Total: totalPrice,
        }
    } else {
        cart[itemID].Quantity += int(quantity)
        cart[itemID].Total += totalPrice
    }

    renderPage(w, "templates/cart.html", cart)
}

func cartData(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json, _ := json.Marshal(&cart)
    w.Write(json)
}
