package App1_Products

import (
    "net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    // logic for handling the home page
    w.Write([]byte("Welcome to the home page!"))
}
