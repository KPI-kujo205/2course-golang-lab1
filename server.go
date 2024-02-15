package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func getTimeHandler(w http.ResponseWriter, _ *http.Request) {
	curTime := map[string]string{"time": time.Now().Format(time.RFC3339)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(curTime)
}

func main() {
	http.HandleFunc("/time", getTimeHandler)
	fmt.Println("Server is running on http://localhost:8795")
	if err := http.ListenAndServe(":8795", nil); err != nil {
		log.Fatal(err)
	}
}
