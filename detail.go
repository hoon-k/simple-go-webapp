package main

import (
    "net/http"
)

func detail(w http.ResponseWriter, r *http.Request) {
    cid := r.URL.Query().Get("cid")
    id := r.URL.Query().Get("id")
    items := inventoryData[cid].Items
    data := items[id]

    renderPage(w, "templates/detail.html", data)
}
