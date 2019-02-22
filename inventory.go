package main

import (
    "net/http"
)

func inventoryHandler(w http.ResponseWriter, r *http.Request) {
    cid := r.URL.Query().Get("cid")
    renderPage(w, "templates/inventory.html", inventoryData[cid])
}
