package main

import (
    "net/http"
)

func detail(w http.ResponseWriter, r *http.Request) {
    cid := r.URL.Query().Get("cid")
    id := r.URL.Query().Get("id")
    items := inventoryData[cid].Items

    var data item

    for i := 0; i < len(items); i++ {
        if (items[i].ID == id) {
            data = items[i]
            break;
        }
    }

    renderPage(w, "templates/detail.html", data)
}
