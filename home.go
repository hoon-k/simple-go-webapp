package main

import (
    "net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
    renderPage(w, "templates/home.html", nil)
}