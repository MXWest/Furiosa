package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/p4tin/Furiosa/app"
)

func main() {
	fmt.Println("Starting Up.")
	http.HandleFunc("/health", health)
	http.ListenAndServe("localhost:50005", nil)
}

func health(w http.ResponseWriter, r *http.Request) {
	response := app.HealthStatusMsg{
		Status:  "OK",
		Message: "All is well in Furiosa world!!!",
	}

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
