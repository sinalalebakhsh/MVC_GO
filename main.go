package main

import (
    "MVC_GO/App1_Products/controllers"
    "MVC_GO/core"
    "net/http"
)

func main() {
    // Create a new server instance
    server := core.NewServer()

    // Define routes
    http.HandleFunc("/", App1_Products.HomeHandler)

    // Start the server
    if err := server.Start(); err != nil {
        panic(err)
    }
}
