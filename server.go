package main

import (
    "fmt"
    "log"
    "net/http"
)

func getTimeHandler(w http.ResponseWriter, r *http.Request) {
    // TODO : implement function for getting current time
}

func main() {
    http.HandleFunc("/time", getTimeHandler)
    fmt.Println("Server is running on http://localhost:8795")
    if err := http.ListenAndServe(":8795", nil); err != nil {
        log.Fatal(err)
    }
}
