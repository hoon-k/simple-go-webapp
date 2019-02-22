package main

import (
    "net/http"
    "strconv"
)

func detail(w http.ResponseWriter, r *http.Request) {
    cid := r.URL.Query().Get("cid")
    id, _ := strconv.ParseInt(r.URL.Query().Get("id"), 0, 64)
    data := inventoryData[cid].Items[id - 1]
    renderPage(w, "templates/detail.html", data)
}
