package main

import (
    "net/http"
    "encoding/json"
)

func inventoryHandler(w http.ResponseWriter, r *http.Request) {
    cid := r.URL.Query().Get("cid")
    renderPage(w, "templates/inventory.html", inventoryData[cid])
}

func inventoryRawData(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json, _ := json.Marshal(allItems)
    w.Write(json)
}
